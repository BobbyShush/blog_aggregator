package cmds

import (
	"bootdev/gator/internal/database"
	"bootdev/gator/internal/config"	
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
