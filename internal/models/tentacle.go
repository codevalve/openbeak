package models

import (
	"context"
	"time"
)

// Tentacle defines the functional interface for all scanner modules.
type Tentacle interface {
	Name() string
	Description() string
	Role() string // Must return Hunter, Reporter, or Beak
	Probe(ctx context.Context, target string) (Result, error)
}

// Result represents a single security finding or discovery status.
type Result struct {
	Target    string    `json:"target"`
	Type      string    `json:"type"`
	Severity  string    `json:"severity"`
	Details   string    `json:"details"`
	Source    string    `json:"source"`
	Timestamp time.Time `json:"timestamp"`
}
