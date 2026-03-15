package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// RequestEditor mutates an HTTP request before sending.
type RequestEditor func(ctx context.Context, req *http.Request) error

// HttpRequestDoer performs HTTP requests.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// ApplyEditorsFn applies any request editors to req.
type ApplyEditorsFn func(ctx context.Context, req *http.Request) error

// Client is a lightweight Omada authorization client.
type Client struct {
	BaseURL      string
	Doer         HttpRequestDoer
	ApplyEditors ApplyEditorsFn
}

// AuthorizeLoginResult contains CSRF/session values from /openapi/authorize/login.
type AuthorizeLoginResult struct {
	CsrfToken string `json:"csrfToken"`
	SessionId string `json:"sessionId"`
}

// TokenResult contains token fields from /openapi/authorize/token.
type TokenResult struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    int64  `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
}

type envelope[T any] struct {
	ErrorCode int    `json:"errorCode"`
	Msg       string `json:"msg"`
	Result    T      `json:"result"`
}

type authorizeLoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authorizeTokenBody struct {
	OmadacId     string `json:"omadacId,omitempty"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// AccessTokenRequestEditor adds Omada's Authorization header format.
func AccessTokenRequestEditor(accessToken string) RequestEditor {
	return func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "AccessToken="+accessToken)
		return nil
	}
}

// AuthorizeLogin calls POST /openapi/authorize/login.
func (c *Client) AuthorizeLogin(ctx context.Context, clientID, omadacID, username, password string) (*AuthorizeLoginResult, error) {
	loginURL, err := c.authorizeURL("/openapi/authorize/login", map[string]string{
		"client_id": clientID,
		"omadac_id": omadacID,
	})
	if err != nil {
		return nil, err
	}

	bodyBytes, err := json.Marshal(authorizeLoginBody{Username: username, Password: password})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, loginURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	if err := c.applyEditors(ctx, req); err != nil {
		return nil, err
	}

	var resp envelope[AuthorizeLoginResult]
	if err := c.doRequest(req, &resp); err != nil {
		return nil, err
	}
	if resp.ErrorCode != 0 {
		return nil, fmt.Errorf("omada authorize login failed: errorCode=%d msg=%s", resp.ErrorCode, resp.Msg)
	}

	return &resp.Result, nil
}

// AuthorizeCode calls POST /openapi/authorize/code.
func (c *Client) AuthorizeCode(ctx context.Context, clientID, omadacID, csrfToken, sessionID string) (string, error) {
	codeURL, err := c.authorizeURL("/openapi/authorize/code", map[string]string{
		"client_id":     clientID,
		"omadac_id":     omadacID,
		"response_type": "code",
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, codeURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Csrf-Token", csrfToken)
	req.Header.Set("Cookie", "TPOMADA_SESSIONID="+sessionID)

	if err := c.applyEditors(ctx, req); err != nil {
		return "", err
	}

	var resp envelope[string]
	if err := c.doRequest(req, &resp); err != nil {
		return "", err
	}
	if resp.ErrorCode != 0 {
		return "", fmt.Errorf("omada authorize code failed: errorCode=%d msg=%s", resp.ErrorCode, resp.Msg)
	}

	return resp.Result, nil
}

// TokenByAuthorizationCode calls POST /openapi/authorize/token in authorization_code mode.
func (c *Client) TokenByAuthorizationCode(ctx context.Context, code, clientID, clientSecret string) (*TokenResult, error) {
	tokenURL, err := c.authorizeURL("/openapi/authorize/token", map[string]string{
		"grant_type": "authorization_code",
		"code":       code,
	})
	if err != nil {
		return nil, err
	}

	return c.tokenRequest(ctx, tokenURL, authorizeTokenBody{ClientId: clientID, ClientSecret: clientSecret})
}

// TokenByClientCredentials calls POST /openapi/authorize/token in client_credentials mode.
func (c *Client) TokenByClientCredentials(ctx context.Context, omadacID, clientID, clientSecret string) (*TokenResult, error) {
	tokenURL, err := c.authorizeURL("/openapi/authorize/token", map[string]string{
		"grant_type": "client_credentials",
	})
	if err != nil {
		return nil, err
	}

	return c.tokenRequest(ctx, tokenURL, authorizeTokenBody{OmadacId: omadacID, ClientId: clientID, ClientSecret: clientSecret})
}

// RefreshToken calls POST /openapi/authorize/token in refresh_token mode.
func (c *Client) RefreshToken(ctx context.Context, refreshToken, clientID, clientSecret string) (*TokenResult, error) {
	tokenURL, err := c.authorizeURL("/openapi/authorize/token", map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	})
	if err != nil {
		return nil, err
	}

	return c.tokenRequest(ctx, tokenURL, authorizeTokenBody{ClientId: clientID, ClientSecret: clientSecret})
}

func (c *Client) authorizeURL(path string, query map[string]string) (string, error) {
	baseURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return "", err
	}

	opPath := path
	if opPath != "" && opPath[0] == '/' {
		opPath = "." + opPath
	}

	requestURL, err := baseURL.Parse(opPath)
	if err != nil {
		return "", err
	}

	values := requestURL.Query()
	for k, v := range query {
		if v != "" {
			values.Set(k, v)
		}
	}
	requestURL.RawQuery = values.Encode()

	return requestURL.String(), nil
}

func (c *Client) tokenRequest(ctx context.Context, tokenURL string, body authorizeTokenBody) (*TokenResult, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	if err := c.applyEditors(ctx, req); err != nil {
		return nil, err
	}

	var resp envelope[TokenResult]
	if err := c.doRequest(req, &resp); err != nil {
		return nil, err
	}
	if resp.ErrorCode != 0 {
		return nil, fmt.Errorf("omada token request failed: errorCode=%d msg=%s", resp.ErrorCode, resp.Msg)
	}

	return &resp.Result, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request) error {
	if c.ApplyEditors == nil {
		return nil
	}
	return c.ApplyEditors(ctx, req)
}

func (c *Client) doRequest(req *http.Request, out any) error {
	if c.Doer == nil {
		return fmt.Errorf("http doer is nil")
	}

	rsp, err := c.Doer.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = rsp.Body.Close() }()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bodyBytes, out); err != nil {
		return fmt.Errorf("decode omada response (%s): %w", rsp.Status, err)
	}

	return nil
}
