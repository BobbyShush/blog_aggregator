package cmds

import (
	"fmt"
	"bootdev/blog_aggregator/internal/database"
	"bootdev/blog_aggregator/internal/config"
)

type State struct {
	Db *database.Queries
	Config *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	M map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	f, exists := c.M[cmd.Name]
	if !exists {
		return fmt.Errorf("Invalid command name")
	}

	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.M[name] = f
}