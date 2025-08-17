package main

import (
	"context"
	"fmt"

	"github.com/L-chaCon/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	newFeed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:     uuid.New(),
			Name:   cmd.Args[0],
			Url:    cmd.Args[1],
			UserID: user.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("not able to create feed. %w", err)
	}

	_, err = s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:     uuid.New(),
			UserID: user.ID,
			FeedID: newFeed.ID,
		})
	if err != nil {
		return fmt.Errorf("not able to follow feed: %w", err)
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
