package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"passme/data"
)

type ConfirmModel struct {
	message        string
	cursor         int
	acceptCallback func()
	rejectCallback func()
	previous       tea.Model
}

// Init implements tea.Model.
func (m ConfirmModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m ConfirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "esc":
			return m.previous, nil

		case "j",
			"k",
			"up",
			"down",
			"h",
			"l",
			"left",
			"right",
			"tab",
			"shift+tab":
			m.cursor = m.cursor*-1 + 1
		case "enter", "y", "n":
			if (msg.String() == "enter" && m.cursor == 0) || msg.String() == "y" {
				m.acceptCallback()
			} else {
				m.rejectCallback()
			}
			sm, ok := m.previous.(data.SyncableModel)
			if ok {
				return sm.Sync(), nil
			}
			return m.previous, nil
		}

	}
	return m, nil
}

// View implements tea.Model.
func (m ConfirmModel) View() string {
	var b strings.Builder
	b.WriteString(m.message + "?\n\n")
	var style0, style1 lipgloss.Style

	if m.cursor == 0 {
		style0 = focusedStyle
		style1 = noStyle
	} else {
		style0 = noStyle
		style1 = focusedStyle
	}

	b.WriteString(style0.Render("Yes"))
	b.WriteString("\n")
	b.WriteString(style1.Render("No"))
	b.WriteString("\n")

	return b.String()
}
