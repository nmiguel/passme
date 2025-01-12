package copy

import (
	"fmt"
	"log"
	"math"
	"os"
	"passme/data"

	"github.com/agnivade/levenshtein"
	clipboard "github.com/atotto/clipboard"
)

func Callback(args []string) {
	if len(args) < 2 {
		fmt.Println("Incorrect number of args, need a target to copy")
		os.Exit(1)
	}
	requested := args[1]

	keys, err := data.GetAllKeys()
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(keys) == 0 {
		fmt.Println("No keys available")
		os.Exit(1)
	}

	bestDistance := math.MaxInt
	bestKey := data.Key{}

	for _, k := range keys {
		d := levenshtein.ComputeDistance(k.Alias, requested)
		if d < bestDistance {
			bestDistance = d
			bestKey = k
		}
	}
	_ = clipboard.WriteAll(bestKey.Token)
	fmt.Println("Successfully copied token with name", bestKey.Alias)
}

