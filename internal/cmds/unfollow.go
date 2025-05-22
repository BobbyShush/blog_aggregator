package cmds

import (
	"fmt"
	"context"
	"bootdev/blog_aggregator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Expecting 1 argument: feed url")
	}

	feedID, err := s.Db.GetFeedID(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("No feed registered under this url")
	}

	params := database.DeleteFollowParams{
		UserID:	user.ID,
		FeedID:	feedID,
	}
	
	err = s.Db.DeleteFollow(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Println("Unfollow completed")
	return nil
}