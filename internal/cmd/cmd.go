package cmd

import (
	"fmt"
	"os"

	mi "github.com/mushtaqx/mi/internal/data/migrations"
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
		fmt.Printf("Command \"%s\" not found\nAvailable commands are: up - down - create\n", args[1])
	}
}
