package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codevalve/openbeak/internal/models"
)

// Msg types
type ResultMsg models.Result
type ScanCompleteMsg struct{}

// Model represents the TUI state.
type Model struct {
	Results  []models.Result
	Status   string
	Scanning bool
	Quitting bool
}

// NewModel initializes the TUI state.
func NewModel() Model {
	return Model{
		Results:  []models.Result{},
		Status:   "Ready to scan the horizon...",
		Scanning: true,
	}
}

// Init sets up the program.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles state transitions.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.Quitting = true
			return m, tea.Quit
		}

	case ResultMsg:
		m.Results = append(m.Results, models.Result(msg))
		m.Status = "Incoming finding detected..."
		return m, nil

	case ScanCompleteMsg:
		m.Scanning = false
		m.Status = "Scan complete. Targets analyzed."
		return m, nil
	}

	return m, nil
}
