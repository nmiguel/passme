package help

import (
	"passme/arguments/flag"
	"fmt"
)

func Callback(args []string, flags []flag.Flag) {
	for _, f := range flags {
		fmt.Println(f.Alias, f.Tooltip)
	}

}
