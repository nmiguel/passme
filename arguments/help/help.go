package help

import (
	"fmt"
	"passme/arguments/flag"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func Callback(args []string, flags []flag.Flag) {
	// for _, f := range flags {
	// 	fmt.Println(f.Alias, f.Tooltip)
	// }
	t := table.New().
		Border(lipgloss.HiddenBorder()).
		Headers("Flag", "Description")

	for _, f := range flags {
		var alias strings.Builder
		for _, a := range f.Alias {
			alias.WriteString(a)
			alias.WriteString(" ")
		}
		t.Row(alias.String(), f.Tooltip)
	}

	fmt.Println(t)

}
