package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGeneratedClientGetMspMailserverStatusBuildsRequestAndAppliesEditors(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("method = %s, want %s", r.Method, http.MethodGet)
		}
		if r.URL.Path != "/openapi/v1/msp/msp-1/account/mail-status" {
			t.Fatalf("path = %s, want %s", r.URL.Path, "/openapi/v1/msp/msp-1/account/mail-status")
		}
		if got := r.Header.Get("X-Global"); got != "1" {
			t.Fatalf("X-Global = %q, want 1", got)
		}
		if got := r.Header.Get("X-Call"); got != "2" {
			t.Fatalf("X-Call = %q, want 2", got)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errorCode":0,"msg":"ok","result":{}}`))
	}))
	defer ts.Close()

	c, err := NewClient(
		ts.URL,
		WithHTTPClient(ts.Client()),
		WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
			req.Header.Set("X-Global", "1")
			return nil
		}),
	)
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	resp, err := c.GetMspMailserverStatus(context.Background(), "msp-1", func(_ context.Context, req *http.Request) error {
		req.Header.Set("X-Call", "2")
		return nil
	})
	if err != nil {
		t.Fatalf("GetMspMailserverStatus returned error: %v", err)
	}
	_ = resp.Body.Close()

	if got := c.Server[len(c.Server)-1:]; got != "/" {
		t.Fatalf("server = %q, expected trailing slash", c.Server)
	}
}

func TestGeneratedClientWithResponsesParsesJSON(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errorCode":0,"msg":"ok","result":{}}`))
	}))
	defer ts.Close()

	c, err := NewClientWithResponses(ts.URL, WithHTTPClient(ts.Client()))
	if err != nil {
		t.Fatalf("NewClientWithResponses returned error: %v", err)
	}

	resp, err := c.GetMspMailserverStatusWithResponse(context.Background(), "msp-2")
	if err != nil {
		t.Fatalf("GetMspMailserverStatusWithResponse returned error: %v", err)
	}
	if resp == nil {
		t.Fatal("response is nil")
	}
	if resp.JSON200 == nil {
		t.Fatal("JSON200 is nil")
	}
	if got, want := resp.StatusCode(), http.StatusOK; got != want {
		t.Fatalf("status = %d, want %d", got, want)
	}
}
