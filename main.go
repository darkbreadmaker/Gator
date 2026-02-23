package main

import (
	"fmt"
	"os"

	"github.com/darkbreadmaker/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg := config.Read()
	appstate := state{config: &cfg}
	cmds := commands{cmds: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	loginStruct := command{
		name: "login",
		args: os.Args[2:],
	}
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%v\n", fmt.Errorf("Not enough arguments, need at least 2"))
		os.Exit(1)
	} else if len(args) < 3 {
		fmt.Printf("%v\n", fmt.Errorf("Please provide a username"))
		os.Exit(1)
	}
	err := cmds.run(&appstate, loginStruct)

	if err != nil {
		fmt.Errorf("Error running command: %s", err)
	}
	cfg = config.Read()
	fmt.Printf("%v", cfg)
}
