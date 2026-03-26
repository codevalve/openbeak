package models

import (
	"context"
	"time"
)

// Tentacle represents a modular scanning unit in OpenBeak.
type Tentacle interface {
	Name() string
	Role() string // Must return Hunter, Reporter, or Beak
	Description() string
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
