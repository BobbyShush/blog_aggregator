package cmds

import (
	"fmt"
	"time"
	"context"
	"github.com/google/uuid"
	"bootdev/gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Expecting 2 arguments: feed name + url") 
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	feed, err := s.Db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	followParams := database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		UserID:		user.ID,
		FeedID:		feed.ID,
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return nil
	}

	fmt.Println("FEED CREATED IN DATABASE")
	fmt.Printf("ID: %v\n", feed.ID)
	fmt.Printf("Created at: %v\n", feed.CreatedAt)
	fmt.Printf("Updated at: %v\n", feed.UpdatedAt)
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("URL: %s\n", feed.Url)
	fmt.Printf("User ID: %v\n", feed.UserID)

	return nil
}