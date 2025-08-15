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

func handlerGetFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("not able to get users. %w", err)
	}
	err = printFeedList(feeds)
	if err != nil {
		return fmt.Errorf("not able to print user list. %w", err)
	}
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * Url:     %v\n", feed.Url)
	fmt.Printf(" * User ID: %v\n", feed.UserID)
	fmt.Println("=====================================")
}

func printFeedUserName(feed database.GetFeedsRow) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * Url:     %v\n", feed.Url)
	fmt.Printf(" * User:    %v\n", feed.UserName)
	fmt.Println("=====================================")
}

func printFeedList(feeds []database.GetFeedsRow) error {
	if len(feeds) == 0 {
		return errors.New("no feeds in the database")
	}
	for _, feed := range feeds {
		printFeedUserName(feed)
	}
	return nil
}
