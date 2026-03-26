package tui

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/codevalve/openbeak/internal/models"
)

// Msg types
type ResultMsg models.Result
type ProgressMsg float64 // 0.0 to 1.0
type ScanCompleteMsg struct{}

// Model represents the TUI state.
type Model struct {
	Results  []models.Result
	Status   string
	Scanning bool
	Quitting bool
	Table    table.Model
	Progress progress.Model
}

// NewModel initializes the TUI state.
func NewModel() Model {
	columns := []table.Column{
		{Title: "Severity", Width: 10},
		{Title: "Target", Width: 25},
		{Title: "Type", Width: 25},
		{Title: "Details", Width: 40},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(MainColor).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	pg := progress.New(progress.WithDefaultGradient())

	return Model{
		Results:  []models.Result{},
		Status:   "Ready to scan the horizon...",
		Scanning: true,
		Table:    t,
		Progress: pg,
	}
}

// Init sets up the program.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles state transitions.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
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

		// Update table rows
		rows := []table.Row{}
		for _, res := range m.Results {
			rows = append(rows, table.Row{
				res.Severity,
				res.Target,
				res.Type,
				res.Details,
			})
		}
		m.Table.SetRows(rows)
		return m, nil

	case ProgressMsg:
		var pgCmd tea.Cmd
		newModel, pgCmd := m.Progress.Update(float64(msg))
		m.Progress = newModel.(progress.Model)
		return m, pgCmd

	case ScanCompleteMsg:
		m.Scanning = false
		m.Status = "Scan complete. Targets analyzed."
		return m, nil
	}

	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}
