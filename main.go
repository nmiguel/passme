package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"passme/arguments"
	"passme/ui"
)

func renderUI() {
	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		renderUI()
	} else {
		arguments.ParseArgs(args)
	}

}
