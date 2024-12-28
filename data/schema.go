package data

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Key struct {
	Alias string
	Token string
}

type SyncableModel interface {
	Sync() tea.Model
}

func (k Key) HideKey() string {
	var displayToken strings.Builder
	for i, c := range k.Token {
		if len(k.Token) > i+4 {
			displayToken.WriteString("*")
		} else {
			displayToken.WriteRune(c)
		}
	}
	return displayToken.String()
}
