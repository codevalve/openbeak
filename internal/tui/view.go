package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View renders the TUI.
func (m Model) View() string {
	if m.Quitting {
		return "Exiting OpenBeak... Stay stealthy.\n"
	}

	var s strings.Builder

	// Header
	s.WriteString(HeaderStyle.Render("OpenBeak | Macroctopus Agentaculum v0.0.1") + "\n\n")

	// Status
	s.WriteString(StatusStyle.Render(fmt.Sprintf("Status: %s", m.Status)) + "\n\n")

	// Findings Table
	if len(m.Results) == 0 {
		if m.Scanning {
			s.WriteString("Scanning for malicious agent deployments...\n")
		} else {
			s.WriteString("No threats found in this sector.\n")
		}
	} else {
		s.WriteString("DETECTED FINDINGS:\n")
		s.WriteString(m.Table.View() + "\n")
	}

	// Progress Bar
	if m.Scanning {
		s.WriteString("\n" + m.Progress.View() + "\n")
	}

	// Footer
	s.WriteString("\n" + lipgloss.NewStyle().Foreground(SubtleColor).Render("Press 'q' or 'esc' to exit.") + " | " + lipgloss.NewStyle().Foreground(MainColor).Render("↑/↓: Navigate") + "\n")

	return s.String()
}
