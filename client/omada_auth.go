package client

import (
	"context"
	"fmt"
	"net/http"
	"omada-go-sdk/client/auth"
)

// OmadaAuthorizeLoginResult contains CSRF/session values from /openapi/authorize/login.
type OmadaAuthorizeLoginResult = auth.AuthorizeLoginResult

// OmadaTokenResult contains token fields from /openapi/authorize/token.
type OmadaTokenResult = auth.TokenResult

// OmadaAccessTokenRequestEditor adds Omada's Authorization header format.
func OmadaAccessTokenRequestEditor(accessToken string) RequestEditorFn {
	editor := auth.AccessTokenRequestEditor(accessToken)
	return func(ctx context.Context, req *http.Request) error {
		return editor(ctx, req)
	}
}

// WithOmadaAccessToken sets Authorization: AccessToken=<token> for requests.
func WithOmadaAccessToken(accessToken string) ClientOption {
	return WithRequestEditorFn(OmadaAccessTokenRequestEditor(accessToken))
}

// OmadaAuthorizeLogin calls POST /openapi/authorize/login.
func (c *Client) OmadaAuthorizeLogin(ctx context.Context, clientID, omadacID, username, password string, reqEditors ...RequestEditorFn) (*OmadaAuthorizeLoginResult, error) {
	return c.omadaAuthClient(reqEditors).AuthorizeLogin(ctx, clientID, omadacID, username, password)
}

// OmadaAuthorizeCode calls POST /openapi/authorize/code.
func (c *Client) OmadaAuthorizeCode(ctx context.Context, clientID, omadacID, csrfToken, sessionID string, reqEditors ...RequestEditorFn) (string, error) {
	return c.omadaAuthClient(reqEditors).AuthorizeCode(ctx, clientID, omadacID, csrfToken, sessionID)
}

// OmadaTokenByAuthorizationCode calls POST /openapi/authorize/token in authorization_code mode.
func (c *Client) OmadaTokenByAuthorizationCode(ctx context.Context, code, clientID, clientSecret string, reqEditors ...RequestEditorFn) (*OmadaTokenResult, error) {
	return c.omadaAuthClient(reqEditors).TokenByAuthorizationCode(ctx, code, clientID, clientSecret)
}

// OmadaTokenByClientCredentials calls POST /openapi/authorize/token in client_credentials mode.
func (c *Client) OmadaTokenByClientCredentials(ctx context.Context, omadacID, clientID, clientSecret string, reqEditors ...RequestEditorFn) (*OmadaTokenResult, error) {
	return c.omadaAuthClient(reqEditors).TokenByClientCredentials(ctx, omadacID, clientID, clientSecret)
}

// OmadaRefreshToken calls POST /openapi/authorize/token in refresh_token mode.
func (c *Client) OmadaRefreshToken(ctx context.Context, refreshToken, clientID, clientSecret string, reqEditors ...RequestEditorFn) (*OmadaTokenResult, error) {
	return c.omadaAuthClient(reqEditors).RefreshToken(ctx, refreshToken, clientID, clientSecret)
}

// OmadaTokenByClientCredentials delegates to the underlying generated client.
func (c *ClientWithResponses) OmadaTokenByClientCredentials(ctx context.Context, omadacID, clientID, clientSecret string, reqEditors ...RequestEditorFn) (*OmadaTokenResult, error) {
	raw, ok := c.ClientInterface.(*Client)
	if !ok {
		return nil, fmt.Errorf("underlying client is not *Client")
	}
	return raw.OmadaTokenByClientCredentials(ctx, omadacID, clientID, clientSecret, reqEditors...)
}

// OmadaRefreshToken delegates to the underlying generated client.
func (c *ClientWithResponses) OmadaRefreshToken(ctx context.Context, refreshToken, clientID, clientSecret string, reqEditors ...RequestEditorFn) (*OmadaTokenResult, error) {
	raw, ok := c.ClientInterface.(*Client)
	if !ok {
		return nil, fmt.Errorf("underlying client is not *Client")
	}
	return raw.OmadaRefreshToken(ctx, refreshToken, clientID, clientSecret, reqEditors...)
}

func (c *Client) omadaAuthClient(reqEditors []RequestEditorFn) *auth.Client {
	return &auth.Client{
		BaseURL: c.Server,
		Doer:    c.Client,
		ApplyEditors: func(ctx context.Context, req *http.Request) error {
			return c.applyEditors(ctx, req, reqEditors)
		},
	}
}
