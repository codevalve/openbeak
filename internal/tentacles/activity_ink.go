package tentacles

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/codevalve/openbeak/internal/models"
)

// ActivityInk writes operational logs to a local file.
type ActivityInk struct {
	FilePath string
	mu       sync.Mutex
}

// Name returns the ink identifier.
func (a *ActivityInk) Name() string {
	return "activity_ink"
}

// Description returns a summary of the ink's purpose.
func (a *ActivityInk) Description() string {
	return "Writes verbose, timestamped operational activity logs to a text file for development and auditing."
}

// Write appends a finding to the log file.
func (a *ActivityInk) Write(ctx context.Context, result models.Result) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	f, err := os.OpenFile(a.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	logEntry := fmt.Sprintf("[%s] [FINDING] [%s] Severity: %s | Target: %s | %s",
		result.Timestamp.Format(time.RFC3339),
		result.Source,
		result.Severity,
		result.Target,
		result.Details,
	)

	_, err = f.WriteString(logEntry + "\n")
	return err
}

// Log appends an activity event to the log file.
func (a *ActivityInk) Log(ctx context.Context, event models.ActivityEvent) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	f, err := os.OpenFile(a.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	logEntry := fmt.Sprintf("[%s] [%s] [%s] %s",
		event.Timestamp.Format(time.RFC3339),
		event.Level,
		event.Component,
		event.Message,
	)
	if event.Target != "" {
		logEntry = fmt.Sprintf("%s (Target: %s)", logEntry, event.Target)
	}

	_, err = f.WriteString(logEntry + "\n")
	return err
}
