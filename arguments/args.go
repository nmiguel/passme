package arguments

import (
	"fmt"
	"os"
	"passme/arguments/copy"
	"slices"
)

type flag struct {
	alias    []string
	tooltip  string
	callback func([]string)
}

var (
	helpFlag = flag{
		alias:    []string{"help", "h", "-h", "--help"},
		tooltip: "Print the help docs",
		callback: helpFunc,
	}
	copyFlag = flag{
		alias:    []string{"copy", "c", "-c", "--copy"},
		tooltip: "Copies the given token name into the clipboard, bypassing the UI",
		callback: copy.Callback,
	}
	flags = []flag{
		copyFlag,
		helpFlag,
	}
)

func helpFunc(args []string) {

}

func ParseArgs(args []string) {
	command := args[0]
	for _, f := range flags {
		if slices.Contains(f.alias, command) {
			f.callback(args)
			return
		}
	}

	fmt.Println("Unknown command or invalid usage")
	os.Exit(1)
}
