package tentacles

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/codevalve/openbeak/internal/models"
)

// ActivityLogger writes operational logs to a local file.
type ActivityLogger struct {
	FilePath string
	mu       sync.Mutex
}

// Name returns the reporter identifier.
func (a *ActivityLogger) Name() string {
	return "activity_logger"
}

// Write is a no-op for findings (reserved for Results).
func (a *ActivityLogger) Write(ctx context.Context, result models.Result) error {
	return nil
}

// Log appends an activity event to the log file.
func (a *ActivityLogger) Log(ctx context.Context, event models.ActivityEvent) error {
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
