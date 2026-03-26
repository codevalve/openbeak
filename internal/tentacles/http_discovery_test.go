package tentacles

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHTTPDiscoveryTentacle_Probe(t *testing.T) {
	// Create a test server to mock responses.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/status" {
			w.Header().Set("X-OpenClaw-Version", "1.2.3")
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.URL.Path == "/health" {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	tent := &HTTPDiscoveryTentacle{Timeout: 1 * time.Second}
	ctx := context.Background()

	// Test case: detection via header
	res, err := tent.Probe(ctx, server.Listener.Addr().String())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if res.Type != "exposed_openclaw_instance" {
		t.Errorf("Expected 'exposed_openclaw_instance', got %s", res.Type)
	}

	// Test case: detection via status code (server will return 200 for /health if we don't have header)
	// We'll mock another server for this to ensure it returns 200 without the header.
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/health" {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server2.Close()

	res2, err := tent.Probe(ctx, server2.Listener.Addr().String())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if res2.Type != "service_discovered" {
		t.Errorf("Expected 'service_discovered' for generic endpoint, got %s", res2.Type)
	}
	if res2.Severity != "Low" {
		t.Errorf("Expected 'Low' severity for generic endpoint, got %s", res2.Severity)
	}
}
