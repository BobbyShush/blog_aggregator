package cmds

import (
	"fmt"
	"time"
	"errors"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"bootdev/blog_aggregator/internal/database"
)

func HandlerRegister(s *State, cmd Command) error{
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Expecting 1 argument: desired username") 
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err == nil {
		return fmt.Errorf("A user with that name already exists")
	} else if !errors.Is(err, sql.ErrNoRows) {
    	// Some other error occurred
    	return err
	}

	params := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
	}

	user, err := s.Db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User created:\nID: %v\nCreated at: %v\nUpdated at: %v\nName: %v\n", 
		user.ID, 
		user.CreatedAt, 
		user.UpdatedAt, 
		user.Name,
	)
	return nil
}