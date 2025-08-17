package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/L-chaCon/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
	return scrapeFeed(s, feed)
}

func scrapeFeed(s *state, feed database.Feed) error {
	updateFeed, err := s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error MarkFeedFetched: %w", err)
	}
	fmt.Printf("updated feed %s at: %v\n", feed.Name, updateFeed.UpdatedAt)

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetchFeed: %w", err)
	}

	for _, item := range feedData.Channel.Item {
		publishAt, err := parseDate(item.PubDate)
		if err != nil {
			return fmt.Errorf("error at publishAt parse: %w", err)
		}

		_, err = s.db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				Title:       sql.NullString{String: item.Title, Valid: true},
				Url:         sql.NullString{String: item.Link, Valid: true},
				Description: sql.NullString{String: item.Description, Valid: true},
				FeedID:      feed.ID,
				PublishedAt: sql.NullTime{Time: publishAt, Valid: true},
			})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code {
				case "23505": // unique_violation
					fmt.Printf("Post already exists (duplicate), skipping: %s\n", item.Title)
				default:
					log.Printf("PostgreSQL error (code: %s) when creating post: %v", pqErr.Code, err)
					return fmt.Errorf("database error: %w", err)
				}
			} else {
				return err
			}
		}
	}
	return nil
}

func parseDate(strinDate string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z, // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC1123,  // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC822Z,  // "02 Jan 06 15:04 -0700"
		time.RFC822,   // "02 Jan 06 15:04 MST"
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, strinDate); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unable to parse date: %s", strinDate)
}
