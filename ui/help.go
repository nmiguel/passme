package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type HelpModel struct {
	previous tea.Model
}

// Init implements tea.Model.
func (m HelpModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m HelpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "q", "esc", "enter":
			return m.previous, nil
		}
	}
	return m, nil
}

// View implements tea.Model.
func (m HelpModel) View() string {
	t := table.New().
		Border(lipgloss.HiddenBorder()).
		Headers("Flag", "Description").
		Rows(
			[]string{
				"a",
				"Add a new key",
			},
			[]string{
				"e",
				"Edit the key under the cursor",
			},
			[]string{
				"q",
				"Return to previous screen",
			},
			[]string{
				"enter\t\t",
				"Copy the token under the cursor",
			},
		)
	return t.String()
}
