package ui

import (
	"log"
	"passme/data"

	clipboard "github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var tableStyle = table.New().
	StyleFunc(func(row, col int) lipgloss.Style {
		style := lipgloss.NewStyle().PaddingLeft(3).PaddingRight(3)
		if row == 0 {
			return style.Align(lipgloss.Center)
		}
		if col == 0 {
			return style.Align(lipgloss.Left)
		}
		return style.Align(lipgloss.Right)
	}).
	Headers("ALIAS", "TOKEN")

type listModel struct {
	keys   []data.Key
	cursor int
}

func InitialModel() listModel {
	keys, err := data.GetAllKeys()
	if err != nil {
		log.Fatalf("could not get keys: %v", err)
	}
	return listModel{
		cursor: 0,
		keys:   keys,
	}
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Sync() tea.Model {
	keys, err := data.GetAllKeys()
	m.keys = keys
	if err != nil {
		log.Fatalf("could not get keys: %v", err)
	}

	if m.cursor >= len(m.keys) {
		m.cursor = len(m.keys) - 1
	}

	return m
}

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.keys)-1 {
				m.cursor++
			}

		case "d":
			return ConfirmModel{
				message: "Delete key: " + m.keys[m.cursor].Alias,
				acceptCallback: func() {
					data.DeleteKey(m.keys[m.cursor].Alias)
					m.keys = append(m.keys[:m.cursor], m.keys[m.cursor+1:]...)
				},
				rejectCallback: func() {},
				previous:       m,
			}, nil

		case "a":
			return NewAddModel(m), nil

		case "e":
			if len(m.keys) > 0 {
				return NewEditModel(m, m.keys[m.cursor]), nil
			}

		case "?":
			return HelpModel{m}, nil

		case "enter", " ":
			// Copies the selected key's token to clipboard
			_ = clipboard.WriteAll(m.keys[m.cursor].Token)
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m listModel) View() string {
	view := titleStyle.Render("Saved keys:") + "\n"

	if len(m.keys) > 0 {
		tableStyle.ClearRows()
		tableStyle.Data(table.NewStringData())
		for i, key := range m.keys {
			prefix := ""
			if i == m.cursor {
				prefix = "> "
			} else {
				prefix = "  "
			}
			tableStyle.Row(prefix+key.Alias, key.HideKey())
		}

		view += tableStyle.Render()
	} else {
		view += "No saved keys"
	}

	view += "\n\nhelp: ?"

	return view
}
