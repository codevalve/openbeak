package models

import "context"

// Reporter defines the interface for sink modules that consume findings and logs.
type Reporter interface {
	Name() string
	Write(ctx context.Context, result Result) error
	Log(ctx context.Context, event ActivityEvent) error
}
