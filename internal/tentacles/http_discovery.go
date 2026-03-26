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

	endpoints := []string{"/", "/health", "/api/v1/status", "/skills/list", "/config"}

	for _, endpoint := range endpoints {
		req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://%s%s", target, endpoint), nil)
		if err != nil {
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// 1. Check for explicit Headers
		if version := resp.Header.Get("X-OpenClaw-Version"); version != "" {
			return models.Result{
				Target:    target,
				Type:      "exposed_openclaw_instance",
				Severity:  "High",
				Details:   fmt.Sprintf("OpenClaw v%s detected via header at %s", version, endpoint),
				Source:    t.Name(),
				Timestamp: time.Now(),
			}, nil
		}

		// 2. Check for endpoint presence
		if resp.StatusCode == http.StatusOK {
			return models.Result{
				Target:    target,
				Type:      "exposed_endpoint",
				Severity:  "Medium",
				Details:   fmt.Sprintf("Unauthenticated access to %s", endpoint),
				Source:    t.Name(),
				Timestamp: time.Now(),
			}, nil
		}
	}

	return models.Result{}, fmt.Errorf("no signature found")
}
