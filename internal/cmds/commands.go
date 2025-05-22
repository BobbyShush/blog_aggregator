package cmds

import (
	"fmt"
)

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