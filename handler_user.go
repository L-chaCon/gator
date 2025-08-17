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
		return fmt.Errorf("not able to find name in database. %w", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("not able to set current user: %w", err)
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
		return errors.New("user already exist")
	}
	newUser, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:   uuid.New(),
			Name: cmd.Args[0],
		},
	)
	if err != nil {
		return fmt.Errorf("not able to create user. %w", err)
	}
	err = s.cfg.SetUser(newUser.Name)
	if err != nil {
		return fmt.Errorf("not able to set user. %w", err)
	}
	printUser(newUser)
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("not able to get users. %w", err)
	}
	err = printUserList(users, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("not able to print user list. %w", err)
	}
	return nil
}
