package ui

import "github.com/charmbracelet/lipgloss"

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	noStyle      = lipgloss.NewStyle()
	pageStyle    = lipgloss.NewStyle().
			Border(lipgloss.BlockBorder())
	titleStyle = lipgloss.NewStyle().
			Width(80).
			Underline(true).
			PaddingTop(1).
			Bold(true)
)
