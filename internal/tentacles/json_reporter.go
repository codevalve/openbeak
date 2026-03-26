package tentacles

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/codevalve/openbeak/internal/models"
)

// JSONReporter writes findings to a local file.
type JSONReporter struct {
	FilePath string
	mu       sync.Mutex
}

// Name returns the reporter identifier.
func (r *JSONReporter) Name() string {
	return "json_reporter"
}

// Write appends a single result to the JSON file.
func (r *JSONReporter) Write(ctx context.Context, result models.Result) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Simplistic implementation: read all, append, write back
	// In a high-volume scenario, we'd use a buffered writer or a live stream
	var results []models.Result
	data, err := os.ReadFile(r.FilePath)
	if err == nil {
		_ = json.Unmarshal(data, &results)
	}

	results = append(results, result)
	out, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.FilePath, out, 0644)
}
