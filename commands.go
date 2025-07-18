package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	callback, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("No handeler for %s command", cmd.Name)
	}
	return callback(s, cmd)
}
