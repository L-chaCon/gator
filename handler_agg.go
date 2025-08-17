package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs: 'h', 'm', 's'>", cmd.Name)
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("not able to parse %s duration", cmd.Args[0])
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error GetNextFeedToFetch: %w", err)
	}

	updateFeed, err := s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error MarkFeedFetched: %w", err)
	}
	fmt.Printf("updated feed %s at: %v\n", feed.Name, updateFeed.UpdatedAt)

	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetchFeed: %w", err)
	}

	for _, item := range fetchedFeed.Channel.Item {
		fmt.Printf(" * Title: %s\n", item.Title)
	}
	fmt.Println("===============================")

	return nil
}
