package cmds

import (
	"fmt"
	"context"
)

func HandlerReset(s *State, cmd Command) error{
	err := s.Db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("Failed database reset. Err: %w", err)
	}
	fmt.Println("Database was successfully reset")
	return nil
}