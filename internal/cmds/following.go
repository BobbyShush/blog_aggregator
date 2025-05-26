package cmds

import (
	"fmt"
	"context"
	"bootdev/gator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Couldn't get feed follows for user %s. Err: %w", user.Name, err)
	}

	fmt.Println("You currently follow the following feeds:")
	for _, ff := range feedFollows {
		fmt.Printf("%s (added by %s)\n", ff.FeedName, ff.CreatorName)
	}

	return nil
}