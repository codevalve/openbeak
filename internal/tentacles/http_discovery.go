package tentacles

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/codevalve/openbeak/internal/models"
)

// HTTPDiscoveryTentacle probes for exposed agent infrastructure.
type HTTPDiscoveryTentacle struct {
	Timeout time.Duration
}

// Name returns the tentacle identifier.
func (t *HTTPDiscoveryTentacle) Name() string {
	return "http_discovery"
}

// Role returns the functional category.
func (t *HTTPDiscoveryTentacle) Role() string {
	return "Hunter"
}

// Description returns a summary of the tentacle's purpose.
func (t *HTTPDiscoveryTentacle) Description() string {
	return "Probes for exposed OpenClaw instances by checking common endpoints (/, /health, /api, /skills) and identifying version headers."
}

// Probe execution logic for a target.
func (t *HTTPDiscoveryTentacle) Probe(ctx context.Context, target string) (models.Result, error) {
	client := &http.Client{
		Timeout: t.Timeout,
	}

	endpoints := []string{"/api/v1/status", "/skills/list", "/config", "/health", "/"}

	for _, endpoint := range endpoints {
		req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://%s%s", target, endpoint), nil)
		if err != nil {
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			continue
		}

		// 1. Check for explicit Headers (High Priority)
		version := resp.Header.Get("X-OpenClaw-Version")
		server := resp.Header.Get("Server")
		statusCode := resp.StatusCode
		_ = resp.Body.Close()

		if version != "" {
			return models.Result{
				Target:    target,
				Type:      "exposed_openclaw_instance",
				Severity:  "High",
				Details:   fmt.Sprintf("OpenClaw v%s detected via header at %s (Server: %s)", version, endpoint, server),
				Source:    t.Name(),
				Timestamp: time.Now(),
			}, nil
		}

		// 2. Check for endpoint presence (Medium vs Low)
		isSpecific := endpoint != "/" && endpoint != "/health"
		if statusCode == http.StatusOK {
			severity := "Low"
			resType := "service_discovered"
			details := fmt.Sprintf("Generic service found at %s (Server: %s)", endpoint, server)

			if isSpecific {
				severity = "Medium"
				resType = "exposed_endpoint"
				details = fmt.Sprintf("Specific OpenClaw endpoint accessible at %s (Server: %s)", endpoint, server)
			}

			return models.Result{
				Target:    target,
				Type:      resType,
				Severity:  severity,
				Details:   details,
				Source:    t.Name(),
				Timestamp: time.Now(),
			}, nil
		}

		// 3. Reconnaissance (Low Priority)
		if statusCode == http.StatusForbidden || statusCode == http.StatusUnauthorized {
			return models.Result{
				Target:    target,
				Type:      "generic_server_found",
				Severity:  "Low",
				Details:   fmt.Sprintf("Server responded with %d at %s (Server: %s)", statusCode, endpoint, server),
				Source:    t.Name(),
				Timestamp: time.Now(),
			}, nil
		}
	}

	return models.Result{}, fmt.Errorf("no signature found")
}
