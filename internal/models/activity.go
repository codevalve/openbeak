package models

import "time"

// ActivityLevel represents the severity of an operational event.
type ActivityLevel string

const (
	Info  ActivityLevel = "INFO"
	Debug ActivityLevel = "DEBUG"
	Warn  ActivityLevel = "WARN"
	Error ActivityLevel = "ERROR"
)

// ActivityEvent represents an internal engine or tentacle operation.
type ActivityEvent struct {
	Timestamp time.Time     `json:"timestamp"`
	Level     ActivityLevel `json:"level"`
	Component string        `json:"component"`
	Message   string        `json:"message"`
	Target    string        `json:"target,omitempty"`
}
