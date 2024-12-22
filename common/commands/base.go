package commands

import (
	"flag"
	"fmt"
	"noisy_neighbour/common/database/commands"
)

type BaseCommand interface {
	Run() error
	GetName() string
}

var Commands = []BaseCommand{&commands.MigrateCommand{}}

func RegisterCommand(command BaseCommand) {
	Commands = append(Commands, command)
}

func RunCommands() {
	for _, command := range Commands {
		flagValue := flag.Bool(command.GetName(), false, fmt.Sprintf("Run %s command", command.GetName()))
		flag.Parse()

		if *flagValue {
			err := command.Run()
			if err != nil {
				fmt.Printf("Error running command %s: %v\n", command.GetName(), err)
			}
			return
		}
	}
}
