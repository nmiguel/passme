package ui

import (
	"passme/data"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type addModel struct {
	previous   tea.Model
	focusIndex int

	aliasInput textinput.Model
	tokenInput textinput.Model
}

func NewAddModel(prev tea.Model) addModel {
	alias := textinput.New()
	alias.Placeholder = "token_name"
	alias.CharLimit = 64
	alias.Focus()
	alias.PromptStyle = focusedStyle
	alias.TextStyle = focusedStyle

	token := textinput.New()
	token.Placeholder = "**********"
	token.CharLimit = 128

	return addModel{
		previous:   prev,
		aliasInput: alias,
		tokenInput: token,
	}
}

func (m addModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m addModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "esc":
			return m.previous, nil

		case "tab", "shift+tab":
			if msg.String() == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}
			if m.focusIndex < 0 {
				m.focusIndex = 1
			} else if m.focusIndex > 1 {
				m.focusIndex = 0
			}
			return updateFocus(m)

		case "enter":
			err := data.InsertKey(m.aliasInput.Value(), m.tokenInput.Value())
			if err != nil {
				// return MessageModel{message: err.Error(), previous: m.previous}, nil
				return m.previous, nil
			}
			sm, ok := m.previous.(data.SyncableModel)
			if ok {
				return sm.Sync(), nil
			}
			return m.previous, nil
		}
	}

	cmds := make([]tea.Cmd, 2)
	m.aliasInput, cmds[0] = m.aliasInput.Update(msg)
	m.tokenInput, cmds[1] = m.tokenInput.Update(msg)
	return m, tea.Batch(cmds...)
}

func (m addModel) View() string {
	var b strings.Builder
	b.WriteString("Adding new key\n\n")

	b.WriteString("Alias:\n")
	b.WriteString(m.aliasInput.View())
	b.WriteString("\n\n")

	b.WriteString("Token:\n")
	b.WriteString(m.tokenInput.View())
	b.WriteString("\n")

	return b.String()
}

func updateFocus(m addModel) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	inputs := []*textinput.Model{&(m.aliasInput), &(m.tokenInput)}

	for i := 0; i < 2; i++ {
		if i == m.focusIndex {
			// Set focused state
			cmd = inputs[i].Focus()
			inputs[i].PromptStyle = focusedStyle
			inputs[i].TextStyle = focusedStyle
			continue
		}
		// Remove focused state
		inputs[i].Blur()
		inputs[i].PromptStyle = noStyle
		inputs[i].TextStyle = noStyle
	}

	return m, tea.Batch(cmd)
}
