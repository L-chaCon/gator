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

	_, err = s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:     uuid.New(),
			UserID: currentUser.ID,
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

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return errors.New("no user found")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return errors.New("no feed found")
	}

	follow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:     uuid.New(),
			UserID: user.ID,
			FeedID: feed.ID,
		})
	if err != nil {
		return fmt.Errorf("not able to follow feed: %w", err)
	}
	printFollow(follow)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return errors.New("not able to found current user")
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return errors.New("not able to found following")
	}
	printFollowing(following)

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

func printFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf(" * ID:           %v\n", feedFollow.ID)
	fmt.Printf(" * User name:    %v\n", feedFollow.UserName)
	fmt.Printf(" * Feed name:    %v\n", feedFollow.FeedName)
	fmt.Println("=====================================")
}

func printFollowing(following []database.GetFeedFollowsForUserRow) error {
	if len(following) == 0 {
		return errors.New("no feeds followed")
	}
	for _, follow := range following {
		fmt.Printf(" * Feed name:   %v\n", follow.Name)
		fmt.Printf(" * URL name:    %v\n", follow.Url)
	}
	return nil
}
