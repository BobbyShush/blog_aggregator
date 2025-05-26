package cmds

import (
	"fmt"
	"context"
	"bootdev/gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Couldn't get user info for current user. Err: %w", err)
		}
		err = handler(s, cmd, user)
		if err != nil {
			return err
		}
		return nil
	}
}
