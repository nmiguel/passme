package add

import (
	"fmt"
	"log"
	"os"
	"passme/data"
)

func Callback(args []string) {
	if len(args) < 3 {
		fmt.Println("Incorrect number of args. Usage: passme add <token_alias> <token>")
		os.Exit(1)
	}
	alias := args[1]
	token := args[2]
	err := data.InsertKey(alias, token)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Successfully added token with name", alias)
}
