package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	err := c.cmds[cmd.name](s, cmd)
	if err != nil {
		fmt.Errorf("Error running command: %s", err)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No arguments, need 1")
	}
	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error setting user: %s", err)
	}
	fmt.Printf("User set")
	return nil
}
