package tui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	MainColor     = lipgloss.Color("#00FFD1") // Neon Teal
	AlertColor    = lipgloss.Color("#FF3131") // High-Alert Red
	SubtleColor   = lipgloss.Color("#555555") // Charcoal
	HighlightColor = lipgloss.Color("#FFFFFF") // White

	// Styles
	HeaderStyle = lipgloss.NewStyle().
			Foreground(MainColor).
			Bold(true).
			Padding(0, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(MainColor)

	FindingStyle = lipgloss.NewStyle().
			Foreground(HighlightColor).
			PaddingLeft(2)

	SeverityHighStyle = lipgloss.NewStyle().
				Foreground(AlertColor).
				Bold(true)

	SeverityMediumStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFA500")).
				Bold(true)

	StatusStyle = lipgloss.NewStyle().
			Foreground(SubtleColor).
			Italic(true)
)
