package auth

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (f roundTripFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

type errReadCloser struct{}

func (errReadCloser) Read(_ []byte) (int, error) {
	return 0, errors.New("read failed")
}

func (errReadCloser) Close() error {
	return nil
}

func TestAccessTokenRequestEditorSetsAuthorizationHeader(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatalf("new request: %v", err)
	}

	editor := AccessTokenRequestEditor("abc123")
	if err := editor(context.Background(), req); err != nil {
		t.Fatalf("editor returned error: %v", err)
	}

	if got, want := req.Header.Get("Authorization"), "AccessToken=abc123"; got != want {
		t.Fatalf("authorization header = %q, want %q", got, want)
	}
}

func TestAuthorizeLoginSendsExpectedRequestAndParsesResult(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/openapi/authorize/login" {
			t.Fatalf("path = %s, want %s", r.URL.Path, "/openapi/authorize/login")
		}
		if got, want := r.URL.Query().Get("client_id"), "client-1"; got != want {
			t.Fatalf("client_id = %q, want %q", got, want)
		}
		if got, want := r.URL.Query().Get("omadac_id"), "omadac-1"; got != want {
			t.Fatalf("omadac_id = %q, want %q", got, want)
		}
		if got := r.Header.Get("Content-Type"); got != "application/json" {
			t.Fatalf("content-type = %q, want application/json", got)
		}
		if got := r.Header.Get("X-Test"); got != "editor-ok" {
			t.Fatalf("X-Test = %q, want editor-ok", got)
		}

		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if got, want := body["username"], "alice"; got != want {
			t.Fatalf("username = %q, want %q", got, want)
		}
		if got, want := body["password"], "secret"; got != want {
			t.Fatalf("password = %q, want %q", got, want)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errorCode":0,"msg":"ok","result":{"csrfToken":"csrf-1","sessionId":"sess-1"}}`))
	}))
	defer ts.Close()

	c := &Client{
		BaseURL: ts.URL,
		Doer:    ts.Client(),
		ApplyEditors: func(_ context.Context, req *http.Request) error {
			req.Header.Set("X-Test", "editor-ok")
			return nil
		},
	}

	result, err := c.AuthorizeLogin(context.Background(), "client-1", "omadac-1", "alice", "secret")
	if err != nil {
		t.Fatalf("AuthorizeLogin returned error: %v", err)
	}
	if result == nil {
		t.Fatal("AuthorizeLogin returned nil result")
	}
	if got, want := result.CsrfToken, "csrf-1"; got != want {
		t.Fatalf("csrfToken = %q, want %q", got, want)
	}
	if got, want := result.SessionId, "sess-1"; got != want {
		t.Fatalf("sessionId = %q, want %q", got, want)
	}
}

func TestTokenByClientCredentialsSendsExpectedRequestAndParsesResult(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/openapi/authorize/token" {
			t.Fatalf("path = %s, want %s", r.URL.Path, "/openapi/authorize/token")
		}
		if got, want := r.URL.Query().Get("grant_type"), "client_credentials"; got != want {
			t.Fatalf("grant_type = %q, want %q", got, want)
		}
		if got := r.Header.Get("Content-Type"); got != "application/json" {
			t.Fatalf("content-type = %q, want application/json", got)
		}

		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if got, want := body["omadacId"], "omadac-2"; got != want {
			t.Fatalf("omadacId = %q, want %q", got, want)
		}
		if got, want := body["client_id"], "client-2"; got != want {
			t.Fatalf("client_id = %q, want %q", got, want)
		}
		if got, want := body["client_secret"], "secret-2"; got != want {
			t.Fatalf("client_secret = %q, want %q", got, want)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errorCode":0,"msg":"ok","result":{"accessToken":"at","tokenType":"Bearer","expiresIn":3600,"refreshToken":"rt"}}`))
	}))
	defer ts.Close()

	c := &Client{BaseURL: ts.URL, Doer: ts.Client()}

	result, err := c.TokenByClientCredentials(context.Background(), "omadac-2", "client-2", "secret-2")
	if err != nil {
		t.Fatalf("TokenByClientCredentials returned error: %v", err)
	}
	if result == nil {
		t.Fatal("TokenByClientCredentials returned nil result")
	}
	if got, want := result.AccessToken, "at"; got != want {
		t.Fatalf("accessToken = %q, want %q", got, want)
	}
	if got, want := result.RefreshToken, "rt"; got != want {
		t.Fatalf("refreshToken = %q, want %q", got, want)
	}
}

func TestAuthorizeCodeSendsExpectedHeadersAndParsesResult(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/openapi/authorize/code" {
			t.Fatalf("path = %s, want %s", r.URL.Path, "/openapi/authorize/code")
		}
		if got, want := r.URL.Query().Get("grant_type"), ""; got != want {
			t.Fatalf("grant_type = %q, want %q", got, want)
		}
		if got, want := r.URL.Query().Get("client_id"), "cid"; got != want {
			t.Fatalf("client_id = %q, want %q", got, want)
		}
		if got, want := r.URL.Query().Get("omadac_id"), "oid"; got != want {
			t.Fatalf("omadac_id = %q, want %q", got, want)
		}
		if got, want := r.URL.Query().Get("response_type"), "code"; got != want {
			t.Fatalf("response_type = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Csrf-Token"), "csrf-x"; got != want {
			t.Fatalf("csrf token = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Cookie"), "TPOMADA_SESSIONID=sess-x"; got != want {
			t.Fatalf("cookie = %q, want %q", got, want)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errorCode":0,"msg":"ok","result":"auth-code-1"}`))
	}))
	defer ts.Close()

	c := &Client{BaseURL: ts.URL, Doer: ts.Client()}

	code, err := c.AuthorizeCode(context.Background(), "cid", "oid", "csrf-x", "sess-x")
	if err != nil {
		t.Fatalf("AuthorizeCode returned error: %v", err)
	}
	if got, want := code, "auth-code-1"; got != want {
		t.Fatalf("code = %q, want %q", got, want)
	}
}

func TestTokenGrantVariantsBuildExpectedQueryAndBody(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		call              func(c *Client) (*TokenResult, error)
		expectedGrantType string
		expectedCode      string
		expectedRefresh   string
		expectedOmadacID  string
	}{
		{
			name: "authorization_code",
			call: func(c *Client) (*TokenResult, error) {
				return c.TokenByAuthorizationCode(context.Background(), "code-1", "cid", "secret")
			},
			expectedGrantType: "authorization_code",
			expectedCode:      "code-1",
		},
		{
			name: "refresh_token",
			call: func(c *Client) (*TokenResult, error) {
				return c.RefreshToken(context.Background(), "rt-1", "cid", "secret")
			},
			expectedGrantType: "refresh_token",
			expectedRefresh:   "rt-1",
		},
		{
			name: "client_credentials",
			call: func(c *Client) (*TokenResult, error) {
				return c.TokenByClientCredentials(context.Background(), "oid-1", "cid", "secret")
			},
			expectedGrantType: "client_credentials",
			expectedOmadacID:  "oid-1",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/openapi/authorize/token" {
					t.Fatalf("path = %s, want %s", r.URL.Path, "/openapi/authorize/token")
				}
				if got, want := r.URL.Query().Get("grant_type"), tc.expectedGrantType; got != want {
					t.Fatalf("grant_type = %q, want %q", got, want)
				}
				if got, want := r.URL.Query().Get("code"), tc.expectedCode; got != want {
					t.Fatalf("code = %q, want %q", got, want)
				}
				if got, want := r.URL.Query().Get("refresh_token"), tc.expectedRefresh; got != want {
					t.Fatalf("refresh_token = %q, want %q", got, want)
				}

				var body map[string]string
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					t.Fatalf("decode body: %v", err)
				}
				if got, want := body["omadacId"], tc.expectedOmadacID; got != want {
					t.Fatalf("omadacId = %q, want %q", got, want)
				}
				if got, want := body["client_id"], "cid"; got != want {
					t.Fatalf("client_id = %q, want %q", got, want)
				}
				if got, want := body["client_secret"], "secret"; got != want {
					t.Fatalf("client_secret = %q, want %q", got, want)
				}

				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"errorCode":0,"msg":"ok","result":{"accessToken":"at","tokenType":"Bearer","expiresIn":1200,"refreshToken":"rt"}}`))
			}))
			defer ts.Close()

			c := &Client{BaseURL: ts.URL, Doer: ts.Client()}
			res, err := tc.call(c)
			if err != nil {
				t.Fatalf("token call returned error: %v", err)
			}
			if res == nil {
				t.Fatal("token result is nil")
			}
		})
	}
}

func TestOmadaAPIReturnedErrorCodeIsPropagated(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/openapi/authorize/code" {
			_, _ = w.Write([]byte(`{"errorCode":1234,"msg":"bad request","result":""}`))
			return
		}
		_, _ = w.Write([]byte(`{"errorCode":1234,"msg":"bad request","result":{}}`))
	}))
	defer ts.Close()

	c := &Client{BaseURL: ts.URL, Doer: ts.Client()}

	if _, err := c.AuthorizeLogin(context.Background(), "cid", "oid", "u", "p"); err == nil || !strings.Contains(err.Error(), "errorCode=1234") {
		t.Fatalf("expected error with errorCode, got: %v", err)
	}
	if _, err := c.AuthorizeCode(context.Background(), "cid", "oid", "csrf", "sid"); err == nil || !strings.Contains(err.Error(), "errorCode=1234") {
		t.Fatalf("expected error with errorCode, got: %v", err)
	}
	if _, err := c.TokenByAuthorizationCode(context.Background(), "code", "cid", "secret"); err == nil || !strings.Contains(err.Error(), "errorCode=1234") {
		t.Fatalf("expected error with errorCode, got: %v", err)
	}
}

func TestInvalidBaseURLReturnsError(t *testing.T) {
	t.Parallel()

	c := &Client{BaseURL: "://bad-url", Doer: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
		return nil, errors.New("should not be called")
	})}

	if _, err := c.AuthorizeLogin(context.Background(), "cid", "oid", "u", "p"); err == nil {
		t.Fatal("expected error for invalid base URL")
	}
}

func TestApplyEditorsHandlesNilAndErrors(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatalf("new request: %v", err)
	}

	c := &Client{}
	if err := c.applyEditors(context.Background(), req); err != nil {
		t.Fatalf("nil apply editors should be no-op: %v", err)
	}

	c.ApplyEditors = func(_ context.Context, _ *http.Request) error {
		return errors.New("editor failed")
	}
	if err := c.applyEditors(context.Background(), req); err == nil || err.Error() != "editor failed" {
		t.Fatalf("expected editor error, got: %v", err)
	}
}

func TestDoRequestErrorBranches(t *testing.T) {
	t.Parallel()

	if err := (&Client{}).doRequest(&http.Request{}, &struct{}{}); err == nil || err.Error() != "http doer is nil" {
		t.Fatalf("expected nil doer error, got: %v", err)
	}

	doerErr := &Client{Doer: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
		return nil, errors.New("network down")
	})}
	if err := doerErr.doRequest(&http.Request{}, &struct{}{}); err == nil || err.Error() != "network down" {
		t.Fatalf("expected doer error, got: %v", err)
	}

	readErr := &Client{Doer: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
		return &http.Response{Status: "200 OK", Body: errReadCloser{}}, nil
	})}
	if err := readErr.doRequest(&http.Request{}, &struct{}{}); err == nil || err.Error() != "read failed" {
		t.Fatalf("expected read error, got: %v", err)
	}

	decodeErr := &Client{Doer: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
		return &http.Response{Status: "200 OK", Body: io.NopCloser(strings.NewReader("not-json"))}, nil
	})}
	if err := decodeErr.doRequest(&http.Request{}, &struct{}{}); err == nil || !strings.Contains(err.Error(), "decode omada response (200 OK)") {
		t.Fatalf("expected decode error, got: %v", err)
	}
}

func TestMethodsReturnErrorWhenApplyEditorsFails(t *testing.T) {
	t.Parallel()

	c := &Client{
		BaseURL: "https://example.com",
		Doer: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return nil, errors.New("doer should not be called")
		}),
		ApplyEditors: func(_ context.Context, _ *http.Request) error {
			return errors.New("editor boom")
		},
	}

	tests := []struct {
		name string
		call func() error
	}{
		{
			name: "authorize login",
			call: func() error {
				_, err := c.AuthorizeLogin(context.Background(), "cid", "oid", "u", "p")
				return err
			},
		},
		{
			name: "authorize code",
			call: func() error {
				_, err := c.AuthorizeCode(context.Background(), "cid", "oid", "csrf", "sid")
				return err
			},
		},
		{
			name: "token auth code",
			call: func() error {
				_, err := c.TokenByAuthorizationCode(context.Background(), "code", "cid", "secret")
				return err
			},
		},
		{
			name: "token client creds",
			call: func() error {
				_, err := c.TokenByClientCredentials(context.Background(), "oid", "cid", "secret")
				return err
			},
		},
		{
			name: "refresh token",
			call: func() error {
				_, err := c.RefreshToken(context.Background(), "rt", "cid", "secret")
				return err
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.call()
			if err == nil || err.Error() != "editor boom" {
				t.Fatalf("expected editor error, got: %v", err)
			}
		})
	}
}

func TestMethodsReturnErrorForInvalidBaseURL(t *testing.T) {
	t.Parallel()

	c := &Client{BaseURL: "://bad-url", Doer: http.DefaultClient}

	tests := []struct {
		name string
		call func() error
	}{
		{
			name: "authorize code",
			call: func() error {
				_, err := c.AuthorizeCode(context.Background(), "cid", "oid", "csrf", "sid")
				return err
			},
		},
		{
			name: "token auth code",
			call: func() error {
				_, err := c.TokenByAuthorizationCode(context.Background(), "code", "cid", "secret")
				return err
			},
		},
		{
			name: "token client creds",
			call: func() error {
				_, err := c.TokenByClientCredentials(context.Background(), "oid", "cid", "secret")
				return err
			},
		},
		{
			name: "refresh token",
			call: func() error {
				_, err := c.RefreshToken(context.Background(), "rt", "cid", "secret")
				return err
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(); err == nil {
				t.Fatal("expected invalid base URL error")
			}
		})
	}
}

func TestTokenRequestWithInvalidURLReturnsError(t *testing.T) {
	t.Parallel()

	c := &Client{Doer: http.DefaultClient}
	_, err := c.tokenRequest(context.Background(), "://bad-url", authorizeTokenBody{ClientId: "cid", ClientSecret: "secret"})
	if err == nil {
		t.Fatal("expected invalid token URL error")
	}
}

func TestAuthorizeURLOmitsEmptyQueryValues(t *testing.T) {
	t.Parallel()

	c := &Client{BaseURL: "https://example.com/base/"}
	u, err := c.authorizeURL("/openapi/authorize/token", map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": "",
	})
	if err != nil {
		t.Fatalf("authorizeURL returned error: %v", err)
	}
	if strings.Contains(u, "refresh_token=") {
		t.Fatalf("authorizeURL should omit empty query values, got: %s", u)
	}
	if !strings.Contains(u, "grant_type=refresh_token") {
		t.Fatalf("missing expected grant_type query in URL: %s", u)
	}
}

func TestAuthorizeURLSupportsRelativeOperationPath(t *testing.T) {
	t.Parallel()

	c := &Client{BaseURL: "https://example.com/base/"}
	u, err := c.authorizeURL("openapi/authorize/login", map[string]string{"client_id": "cid"})
	if err != nil {
		t.Fatalf("authorizeURL returned error: %v", err)
	}
	if !strings.Contains(u, "/base/openapi/authorize/login") {
		t.Fatalf("unexpected relative path handling: %s", u)
	}
}

func TestMethodsPropagateDecodeErrorsFromDoRequest(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not-json"))
	}))
	defer ts.Close()

	c := &Client{BaseURL: ts.URL, Doer: ts.Client()}

	tests := []struct {
		name string
		call func() error
	}{
		{
			name: "authorize login",
			call: func() error {
				_, err := c.AuthorizeLogin(context.Background(), "cid", "oid", "u", "p")
				return err
			},
		},
		{
			name: "authorize code",
			call: func() error {
				_, err := c.AuthorizeCode(context.Background(), "cid", "oid", "csrf", "sid")
				return err
			},
		},
		{
			name: "token auth code",
			call: func() error {
				_, err := c.TokenByAuthorizationCode(context.Background(), "code", "cid", "secret")
				return err
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.call()
			if err == nil || !strings.Contains(err.Error(), "decode omada response") {
				t.Fatalf("expected decode error, got: %v", err)
			}
		})
	}
}
