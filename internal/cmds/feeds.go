package cmds

import (
	"fmt"
	"context"
)

func HandlerFeeds(s *State, cmd Command) error{
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for i, f := range feeds {
		fmt.Printf("----------FEED %d----------\n", i+1)
		fmt.Printf("ID: %v\n", f.ID)
		fmt.Printf("Name: %v\n", f.Name)
		fmt.Printf("URL: %v\n", f.Url)
		fmt.Printf("User name: %v\n", f.CreatorName.String)
	}
	return nil
}