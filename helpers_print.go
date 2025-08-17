package main

import (
	"errors"
	"fmt"

	"github.com/L-chaCon/gator/internal/database"
)

// PRINTER FOR POSTS
func printPost(post database.Post) {
	fmt.Printf(" * Title:       %s\n", post.Title.String)
	fmt.Printf(" * Description: %s\n", post.Description.String)
	fmt.Printf(" * publishAt:   %s\n", post.PublishedAt.Time)
	fmt.Printf(" * Url:         %s\n", post.Url.String)
	fmt.Println("==========================")
}

// PRINTER FOR USERS
func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}

func printUserList(users []database.User, currentUser string) error {
	if len(users) == 0 {
		return errors.New("no users in the database")
	}
	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}

// PRINTERS FOR FEEDS
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
