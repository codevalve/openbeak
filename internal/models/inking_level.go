package models

import "strings"

// InkingLevel defines how much detail the engine exports to inks.
type InkingLevel string

const (
	InkingStealth  InkingLevel = "stealth"  // Only high-severity (confirmed Claw)
	InkingTactical InkingLevel = "tactical" // Standard findings (High/Medium)
	InkingVerbose  InkingLevel = "verbose"  // All probes and system events (High/Medium/Low)
)

// ParseInkingLevel converts a string to an InkingLevel, defaulting to tactical.
func ParseInkingLevel(s string) InkingLevel {
	switch strings.ToLower(s) {
	case "stealth":
		return InkingStealth
	case "verbose":
		return InkingVerbose
	case "tactical":
		return InkingTactical
	default:
		return InkingStealth // OOB Default is Stealth
	}
}
