package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"passme/ui"
)

func renderUI()  {
	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//Main window func
	//Table with columns: alias, value
	//Add new entry
	//Remove entry
	//Edit entry
	//Navigate, highlight current selected

	//Calling just passme will open the GUI
	renderUI()

}
