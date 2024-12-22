package main

import (
	"noisy_neighbour/common/commands"
	"noisy_neighbour/common/database"
)

func main() {
	commands.RunCommands()

	_, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

}
