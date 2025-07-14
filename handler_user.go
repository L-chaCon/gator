package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/L-chaCon/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Not able to find name in database. %w", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Not able to set current user: %w", err)
	}

	fmt.Printf("You have been login with %s\n", cmd.Args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err == nil {
		return errors.New("User already exist")
	}
	newUser, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:   uuid.New(),
			Name: cmd.Args[0],
		},
	)
	if err != nil {
		return fmt.Errorf("Not able to create user. %w", err)
	}
	err = s.cfg.SetUser(newUser.Name)
	if err != nil {
		return fmt.Errorf("Not able to set user. %w", err)
	}
	fmt.Printf("This is the new User %+v\n", newUser)

	return nil
}
