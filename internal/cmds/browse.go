package cmds

import (
	"fmt"
	"strconv"
	"context"
	"bootdev/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit32 := int32(2)
	if len(cmd.Args) > 0 {
		limit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("Expected 1 argument: limit. Err: %v", err)
		}
		limit32 = int32(limit)
	}
	
	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: limit32,
	}
	
	posts, err := s.Db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}

	for i, post := range posts {
		fmt.Printf("-------- POST %d--------\n", i+1)
		displayPost(post)
	}

	return nil
}

func displayPost(post database.Post) {
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Url: %s\n", post.Url)
	if post.Description.Valid {
		fmt.Printf("Description: %s\n", post.Description.String)
	}else {
		fmt.Println("Description: (None)")
	}
	if post.PublishedAt.Valid {
		fmt.Printf("Published at: %v\n", post.PublishedAt.Time)
	}else {
		fmt.Println("Published at: (None)")
	}
}