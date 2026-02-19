package main

import (
	"fmt"
	"os"

	"github.com/darkbreadmaker/internal/config"
)

func main() {
	state := state{Config: config.Read()}
	cmds := commands{cmds: make(map[string]func(*state, command) error)}
	err := cmds.register("login", handlerLogin)
	if err != nil {
		fmt.Errorf("Error registering command: %s", err)
	}
	args := os.Args
	if len(args) < 2 {
		fmt.Errorf("Not enough arguments, need at least 2")
	}
	var argsSlice []string
	for i := 1; i < len(args); i++ {
		argsSlice = append(argsSlice, args[i])
	}
	err = cmds.run(state, argsSlice[0])
	if err != nil {
		fmt.Errorf("Error running command: %s", err)
>>>>>>> 9c781bf (Fixed bugs and added partial command functionality)
	}
	cfg = config.Read()
	fmt.Printf("%v", cfg)
}
