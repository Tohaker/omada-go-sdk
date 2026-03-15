package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWithOmadaAccessTokenAddsExpectedAuthorizationHeader(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("Authorization"), "AccessToken=tok-1"; got != want {
			t.Fatalf("authorization = %q, want %q", got, want)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errorCode":0,"msg":"ok","result":{"accessToken":"at","tokenType":"Bearer","expiresIn":1,"refreshToken":"rt"}}`))
	}))
	defer ts.Close()

	c, err := NewClient(ts.URL, WithHTTPClient(ts.Client()), WithOmadaAccessToken("tok-1"))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	res, err := c.OmadaTokenByClientCredentials(context.Background(), "oid", "cid", "secret")
	if err != nil {
		t.Fatalf("OmadaTokenByClientCredentials returned error: %v", err)
	}
	if res == nil {
		t.Fatal("result is nil")
	}
}

func TestOmadaClientWithResponsesDelegatesToUnderlyingClient(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/openapi/authorize/token" {
			t.Fatalf("path = %s, want %s", r.URL.Path, "/openapi/authorize/token")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errorCode":0,"msg":"ok","result":{"accessToken":"at","tokenType":"Bearer","expiresIn":1,"refreshToken":"rt"}}`))
	}))
	defer ts.Close()

	raw, err := NewClient(ts.URL, WithHTTPClient(ts.Client()))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	cwr := &ClientWithResponses{ClientInterface: raw}
	res, err := cwr.OmadaTokenByClientCredentials(context.Background(), "oid", "cid", "secret")
	if err != nil {
		t.Fatalf("OmadaTokenByClientCredentials returned error: %v", err)
	}
	if res == nil {
		t.Fatal("result is nil")
	}

	res, err = cwr.OmadaRefreshToken(context.Background(), "rt-old", "cid", "secret")
	if err != nil {
		t.Fatalf("OmadaRefreshToken returned error: %v", err)
	}
	if res == nil {
		t.Fatal("result is nil")
	}
}

func TestOmadaClientWithResponsesRejectsNonConcreteClient(t *testing.T) {
	t.Parallel()

	cwr := &ClientWithResponses{}

	if _, err := cwr.OmadaTokenByClientCredentials(context.Background(), "oid", "cid", "secret"); err == nil || !strings.Contains(err.Error(), "underlying client is not *Client") {
		t.Fatalf("expected type assertion error, got: %v", err)
	}
	if _, err := cwr.OmadaRefreshToken(context.Background(), "rt", "cid", "secret"); err == nil || !strings.Contains(err.Error(), "underlying client is not *Client") {
		t.Fatalf("expected type assertion error, got: %v", err)
	}
}
