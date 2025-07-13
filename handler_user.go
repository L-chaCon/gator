package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Not able to set current user: %w", err)
	}

	fmt.Printf("You have been login with %s\n", cmd.Args[0])
	return nil
}
