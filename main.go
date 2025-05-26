package main

import (
	"fmt"
	"os"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"bootdev/gator/internal/database"
	"bootdev/gator/internal/config"
	"bootdev/gator/internal/cmds"
)

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read config file. Err: %v", err)
	}

	state := cmds.State{ Config: &cfg }
	
	db, err := sql.Open("postgres", state.Config.DbUrl)
	if err != nil {
		log.Fatalf("Failed to open database. Err: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	state.Db = dbQueries

	commands := cmds.Commands{ 
		M: map[string]func(*cmds.State, cmds.Command) error{},
	}
	commands.Register("login", cmds.HandlerLogin)
	commands.Register("register", cmds.HandlerRegister)
	commands.Register("reset", cmds.HandlerReset)
	commands.Register("users", cmds.HandlerUsers)
	commands.Register("agg", cmds.HandlerAgg)
	commands.Register("addfeed", cmds.MiddlewareLoggedIn(cmds.HandlerAddFeed))
	commands.Register("feeds", cmds.HandlerFeeds)
	commands.Register("follow", cmds.MiddlewareLoggedIn(cmds.HandlerFollow))
	commands.Register("following", cmds.MiddlewareLoggedIn(cmds.HandlerFollowing))
	commands.Register("unfollow", cmds.MiddlewareLoggedIn(cmds.HandlerUnfollow))
	commands.Register("browse", cmds.MiddlewareLoggedIn(cmds.HandlerBrowse))

	lauchArgs := os.Args
	if len(lauchArgs) < 2 {
		log.Fatalf("Please enter a valid command. Format: gator <command name> <arguments>")
	}

	command := cmds.Command{
		Name: lauchArgs[1],
		Args: lauchArgs[2:],
	}

	err = commands.Run(&state, command)
	if err != nil {
		log.Fatal(err)
	}
}