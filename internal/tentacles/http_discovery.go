package tentacles

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/codevalve/openbeak/internal/models"
)

// HTTPDiscoveryTentacle probes for exposed OpenClaw API signatures.
type HTTPDiscoveryTentacle struct {
	Timeout time.Duration
}

// Name returns the tentacle identifier.
func (t *HTTPDiscoveryTentacle) Name() string {
	return "http_discovery"
}

// Description returns what this tentacle does.
func (t *HTTPDiscoveryTentacle) Description() string {
	return "Probes for exposed OpenClaw API signatures and health endpoints."
}

// Role returns the functional category.
func (t *HTTPDiscoveryTentacle) Role() string {
	return "Hunter"
}

// Probe execution logic for a target.
func (t *HTTPDiscoveryTentacle) Probe(ctx context.Context, target string) (models.Result, error) {
	client := &http.Client{
		Timeout: t.Timeout,
	}

	// Example: Probing the /health endpoint
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://%s/health", target), nil)
	if err != nil {
		return models.Result{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return models.Result{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return models.Result{
			Target:    target,
			Type:      "exposed_health_endpoint",
			Severity:  "Medium",
			Details:   "Health endpoint responded with 200 OK",
			Source:    t.Name(),
			Timestamp: time.Now(),
		}, nil
	}

	return models.Result{}, fmt.Errorf("no signature found")
}
