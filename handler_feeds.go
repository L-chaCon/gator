package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/L-chaCon/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return errors.New("user not found")
	}

	newFeed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:     uuid.New(),
			Name:   cmd.Args[0],
			Url:    cmd.Args[1],
			UserID: currentUser.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("not able to create feed. %w", err)
	}
	printFeed(newFeed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * Url:     %v\n", feed.Url)
	fmt.Printf(" * User ID: %v\n", feed.UserID)
}
