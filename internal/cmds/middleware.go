package cmds

import (
	"context"
	"bootdev/blog_aggregator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}
		err = handler(s, cmd, user)
		if err != nil {
			return err
		}
		return nil
	}
}
