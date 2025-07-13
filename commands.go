package main

import (
	"fmt"
)

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	callback, ok := c.commands[cmd.name]
	if !ok {
		return fmt.Errorf("No handeler for %s command", cmd.name)
	}

	return callback(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Username is required.")
	}
	err := s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Printf("You have been login with %s\n", cmd.arguments[0])
	return nil
}
