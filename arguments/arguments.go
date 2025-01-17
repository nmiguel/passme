package arguments

import (
	"fmt"
	"os"
	"passme/arguments/add"
	"passme/arguments/copy"
	"passme/arguments/flag"
	"passme/arguments/help"
	"slices"
)

var (
	flags     []flag.Flag
	helpFlag  flag.Flag
	copyFlag  flag.Flag
	addFlag  flag.Flag
)

func init() {
	helpFlag = flag.Flag{
		Alias:   []string{"help", "h"},
		Tooltip: "Print the help docs",
		Callback: func(args []string) {
			help.Callback(args, flags)
		},
	}

	copyFlag = flag.Flag{
		Alias:    []string{"copy", "c"},
		Tooltip: "Copies the given token name into the clipboard.\nUsage: passme copy <token_alias>",
		Callback: copy.Callback,
	}

	addFlag = flag.Flag{
		Alias:    []string{"add", "a"},
		Tooltip: "Add a new token directly.\nUsage: passme add <token_alias> <token>",
		Callback: add.Callback,
	}

	flags = []flag.Flag{copyFlag, helpFlag, addFlag}
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
