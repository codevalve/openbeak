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

	// Findings
	if len(m.Results) == 0 {
		if m.Scanning {
			s.WriteString("Scanning for malicious agent deployments...\n")
		} else {
			s.WriteString("No threats found in this sector.\n")
		}
	} else {
		s.WriteString("DETECTED FINDINGS:\n")
		for _, res := range m.Results {
			sevStyle := SeverityMediumStyle
			if res.Severity == "High" {
				sevStyle = SeverityHighStyle
			}

			s.WriteString(fmt.Sprintf("%s [%s] %s -> %s\n",
				sevStyle.Render(fmt.Sprintf("[%s]", strings.ToUpper(res.Severity))),
				lipgloss.NewStyle().Foreground(SubtleColor).Render(res.Target),
				res.Type,
				res.Details,
			))
		}
	}

	s.WriteString("\n" + lipgloss.NewStyle().Foreground(SubtleColor).Render("Press 'q' or 'esc' to exit.") + "\n")

	return s.String()
}
