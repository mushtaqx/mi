package cmd

import (
	"fmt"
	"os"

	mi "github.com/mushtaqx/migo/data/migrations"
)

// Execute root command, log and exit on error
func Execute() {
	if len(os.Args) < 3 {
		fmt.Println("Missing arguments")
		return
	}
	if os.Args[1] == "mi" {
		handle(os.Args)
	}
}

func handle(args []string) {
	switch args[2] {
	case "up":
		mi.Up()
	case "down":
		mi.Down()
	case "create":
		mi.Create(args[3])
	default:
		fmt.Printf("command \"%s\" not found\n", args[1])
	}
}
