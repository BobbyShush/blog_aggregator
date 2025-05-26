package cmds

import (
	"fmt"
	"context"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Expecting 1 argument: username") 
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("The user '%v' hasn't been registered yet", cmd.Args[0])
	}

	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Failed login. Err: %w", err)
	}
	fmt.Printf("User set to: %s\n", cmd.Args[0])
	return nil
}