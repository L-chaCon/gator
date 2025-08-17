package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/L-chaCon/gator/internal/database"
)

func headerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2

	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <int>", cmd.Name)
	}

	if len(cmd.Args) == 1 {
		val, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil || val <= 0 {
			return fmt.Errorf("invalid limit '%s': must be a positive number", cmd.Args[0])
		}
		limit = int32(val)
	}

	posts, err := s.db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  limit,
		},
	)
	if err != nil {
		return fmt.Errorf("not able to get post: %w", err)
	}
	for _, post := range posts {
		printPost(post)
	}
	return nil
}
