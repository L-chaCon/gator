package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/L-chaCon/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s", cmd.Name)
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

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return errors.New("not able to found following")
	}

	if len(following) == 0 {
		fmt.Println("no feeds followed by this user")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, follow := range following {
		fmt.Printf(" * Feed name:    %s\n", follow.Name.String)
	}
	fmt.Println("=====================================")
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return errors.New("not able to find feed")
	}

	_, err = s.db.UnfollowForUser(
		context.Background(),
		database.UnfollowForUserParams{
			UserID: user.ID,
			FeedID: feed.ID,
		})
	if err != nil {
		return errors.New("not able to unfollow")
	}
	fmt.Printf("User: %s has unfollow %s", user.Name, feed.Url)
	return nil
}

func printFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf(" * ID:           %v\n", feedFollow.ID)
	fmt.Printf(" * User name:    %v\n", feedFollow.UserName)
	fmt.Printf(" * Feed name:    %v\n", feedFollow.FeedName)
	fmt.Println("=====================================")
}
