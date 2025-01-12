package arguments

import (
	"fmt"
	"os"
	"passme/arguments/copy"
	"passme/arguments/help"
	"passme/arguments/flag"
	"slices"
)

var (
	flags     []flag.Flag
	helpFlag  flag.Flag
	copyFlag  flag.Flag
)

func init() {
	helpFlag = flag.Flag{
		Alias:   []string{"help", "h", "-h", "--help"},
		Tooltip: "Print the help docs",
		Callback: func(args []string) {
			help.Callback(args, flags)
		},
	}

	copyFlag = flag.Flag{
		Alias:    []string{"copy", "c", "-c", "--copy"},
		Tooltip: "Copies the given token name into the clipboard, bypassing the UI",
		Callback: copy.Callback,
	}

	flags = []flag.Flag{copyFlag, helpFlag}
}

func GetAvailableFlags() []flag.Flag {
	return flags
}


func ParseArgs(args []string) {
	command := args[0]
	for _, f := range flags {
		if slices.Contains(f.Alias, command) {
			f.Callback(args)
			return
		}
	}

	fmt.Println("Unknown command or invalid usage")
	os.Exit(1)
}
