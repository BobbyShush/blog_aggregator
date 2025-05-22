package cmds

import (
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
