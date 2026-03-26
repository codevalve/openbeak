package models

import "context"

// Ink defines the interface for sink modules that consume findings and logs.
type Ink interface {
	Name() string
	Description() string
	Write(ctx context.Context, result Result) error
	Log(ctx context.Context, event ActivityEvent) error
}
