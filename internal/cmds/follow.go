package cmds

import (
	"fmt"
	"time"
	"context"
	"github.com/google/uuid"
	"bootdev/gator/internal/database"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Expecting 1 argument: feed url") 
	}

	feedID, err := s.Db.GetFeedID(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Couldn't get feed ID for url %s. Err: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID: 		uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		UserID:		user.ID,
		FeedID:		feedID,
	}

	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Couldn't create follow. Err: %w", err)
	}
	
	fmt.Printf("User %s now follows feed %s\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}