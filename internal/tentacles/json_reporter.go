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

// Description returns a summary of the reporter's purpose.
func (r *JSONReporter) Description() string {
	return "Exports all discovery findings to a structured JSON file for automation/SIEM integration."
}

// Write appends a single result to the JSON file.
func (r *JSONReporter) Write(ctx context.Context, result models.Result) error {
	r.mu.Lock()
	defer r.mu.Unlock()

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

// Log is currently not implemented for the JSON results reporter.
func (r *JSONReporter) Log(ctx context.Context, event models.ActivityEvent) error {
	return nil
}
