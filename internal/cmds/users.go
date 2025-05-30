package cmds

import (
	"fmt"
	"context"
)

func HandlerUsers(s *State, cmd Command) error{
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to get users in the database. Err: %w", err)
	}

	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}