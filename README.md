# blog_aggregator
My implementation of the boot.dev guided project "Build a Blog Aggregator in Go"

# User guide
## Description
Gator (short for aggregator) is a CLI aggregator that fetches RSS feeds from the web, stores the feeds and the posts it contains into a database and then displays the posts in the CLI. The application stores informations pertaining to the user, meaning that different users can follow different feeds and be presented only with the choices of feeds matching their own preferences.

## How to use it?
To be able to use Gator, the user needs to install the following dependencies:
- Go 1.24.3 or higher
- Postgres 15 or higher
Once the dependencies are installed, the user can install gator and start using it directly with the different commands listed below.
It is assumed the user has knowledge on how to use a CLI.

### Install Go
The simplest way to install Go is to use [Webi](https://webinstall.dev/golang/).

Go can also be installed from the [Go website](https://go.dev/doc/install). Package managers (like Homebrew on MacOs, apt on Unix systems) also have means to install it easily and information for those commands in the Go website as well.
A successful installation can be verified by using the command:
`go version`
For example my output with this command looks like this:
`go version go1.24.3 linux/amd64`

### Install Postgres
The simplest way to install Postgres is to use package managers.
MacOS:
`brew install postgresql@15`
Linux/WSL:
```bash
sudo apt update
sudo apt install postgresql-contrib
```
Follow the information from [the PostgreSQL website](https://www.postgresql.org/download/windows/) to install on Windows (withouth WSL).

- To confirm the install, simply use:
`psql --version`
- Linux/WSL users need to set up a new system password for Postgres:
`sudo passwd postgres`
- Then the server needs to be started in the background:
Mac: `brew services start postgresql@15`
Linux/WSL: `sudo service postgresql start`
- Enter the psql shell:
Mac: `psql postgres`
Linux: `sudo -u postgres psql`
- A new prompt will be displayed:
`postgres=#`
- Create a new database:
`CREATE DATABASE gator;`
- Connect to the database:
`\c gator`
- The prompt with change to this:
`gator=#`
- (Linux/WSL user) Set up and in-app password:
`ALTER USER postgres PASSWORD '<password>';`
- Type `exit` to go back to the shell.

### Install Gator and finish setting up the database
- Clone the repository from Github
`git clone https://github.com/BobbyShush/blog_aggregator`
- From the base of the repository, use the go install command.
This will ensure the app can be used with the gator command from anywhere.
`go install .`

- Install Goose:
`go install github.com/pressly/goose/v3/cmd/goose@latest`
- Run up the migration. Use this command from the base of the repository with the proper credentials.
- - MacOS - System username, no password
- - Linux/WSL - Username "postgres", in-app password (set earlier)
```bash
goose -dir ./sql/schema postgres postgres://<username>:<password>@localhost:5432/gator up
```

### Config setup
In the home directory, create a file named .gatorconfig.json and paste this content inside (adapt the username and password accordingly):
```json
{"db_url":"postgres://<username>:<password>@localhost:5432/gator?sslmode=disable","current_user_name":""}
```

## Command list
### Register
```bash
gator register <username>
```

Register a user to the database. A logged in user is necessery for adding, following, unfollowing, aggregating and browsing feeds. A newly registered user is automatically logged in as the active user.

Upon successfully registering a user, a message like this one will be displayed:
```
User created:
ID: <uuid>
Created at: 2025-05-26 11:53:49.879511 +0000 +0000
Updated at: 2025-05-26 11:53:49.879511 +0000 +0000
Name: <username>
```

### Login
```bash
gator login <username>
```

Log in a different user than the one currently saved in the config file.
Note : There is no logout option, the active user is the last one that was logged into the config file.

Upon a successful login, a message like this one will be displayed:
```
User set to: <username>
```

### Users
```bash
gator users
```
Lists all the users registered in the database, indicating which user is active.
Example:
```
* bobby
* bernadette
* sheldon (current)
* leonard
* penny
```

### Reset (for testing only)
```bash
gator reset
```
Entirely resets the database, removing all users, feeds, follows and posts.

Upon a successful reset, a message like this one will be displayed:
```
Database was successfully reset
```

### Add feed
```bash
gator addfeed <feed name> <feed url>
```
Adds a feed to the database and adds it in the logged in user's follows.
Note: A name can contain multiple words if they are contained between quotes. Example: "Hacker News RSS".

Upon successfully adding a feed, a message like this one will be displayed:
```
FEED CREATED IN DATABASE
ID: <uuid>
Created at: 2025-05-26 12:06:16.949058 +0000 +0000
Updated at: 2025-05-26 12:06:16.949058 +0000 +0000
Name: <feed name>
URL: <feed url>
User ID: <uuid of the current user>
```

### Feeds
```bash
gator feeds
```
Lists all the feeds currently in the database.
Here is an example output:
```
----------FEED 1----------
ID: <uuid>
Name: <feed name>
URL: <feed url>
User name: <username of the user that added the feed>
```

### Follow
```bash
gator follow <url>
```
Links a user to a feed so that only feeds matching the user's preferences are aggregated and only their posts are displayed when browsed.

Upon successfully registering a follow, a message like this one will be displayed:
```
User linda now follows feed Hacker News RSS
```

### Following
```bash
gator following
```
Lists all the feeds the current user is following.
Example output:
```
You currently follow the following feeds:
<feed name> (added by <username>)
```

### Unfollow
```bash
gator unfollow <url>
```
Removes the link between a user and a feed, so they won't be displayed in browse results.
Upon successfully unfollowing a feed, a message like this one will be displayer:
```
Unfollow completed
```

### Agg
```bash
gator agg <time between requests>
```
Gathers all the posts from the feeds followed by the user and stores them in the database.
The "time between requests" argument determine at what interval this action is executed. The argument is expected as a duration string (example: 2h for 2 hours, 2m for 2 minutes, 2s for 2 seconds). The information of one feed (the feed that hasn't been fetched from in the longest time) is fetched at every interval specified. *This action is executed in a loop until the user stops it by using CTRL + C.*

> WARNING: Using a short interval can result it DOSing the websites from which the feeds are fetched. Caution is advised.

While this command is being executed, the terminal will display this message:
```
Collecting feeds every <time between requests>
```
Gator can be used on a second terminal while the agg command is being executed, or the user can leave it running until satisfied and *interrupt it with CTRL + C* to continue using gator in the same terminal.

### Browse
```bash
gator browse <limit>
```
Displays a number of posts <= to the provided limit argument. The posts are displayed from the most recent to the oldest.
Note: Failing to provide a limit argument will automatically set it to 2.

Example of display:
```
-------- POST 1--------
Title: Genius
Url: https://www.americanscientist.org/article/hidden-genius
Description:
<p>Article URL: <a href="https://www.americanscientist.org/article/hidden-genius">https://www.americanscientist.org/article/hidden-genius</a></p>
<p>Comments URL: <a href="https://news.ycombinator.com/item?id=44098886">https://news.ycombinator.com/item?id=44098886</a></p>
<p>Points: 1</p>
<p># Comments: 0</p>

Published at: 2025-05-26 16:24:01 +0000 +0000
-------- POST 2--------
Title: How to Become a Student of Quantonics
Url: https://www.quantonics.com/How_to_Become_A_Student_of_Quantonics.html
Description:
<p>Article URL: <a href="https://www.quantonics.com/How_to_Become_A_Student_of_Quantonics.html">https://www.quantonics.com/How_to_Become_A_Student_of_Quantonics.html</a></p>
<p>Comments URL: <a href="https://news.ycombinator.com/item?id=44098877">https://news.ycombinator.com/item?id=44098877</a></p>
<p>Points: 1</p>
<p># Comments: 0</p>

Published at: 2025-05-26 16:23:02 +0000 +0000
```

# Dev guide
## Commands and how they are implemented
Commands handle communication with the database through SQL queries (defined in the database package) and RSS fetching through HTTP requests (defined in the rss package).
They are implemented individually in the cmds package. They can have two types of signatures:
- Regular signature
```go
func HandlerCmdName(s *State, cmd Command) error
```
- Logged-in signature
```go
func HandlerCmdName(s *State, cmd Command, user database.User) error
```

Commands are registered in the "commands" map instantiated in main.go through the commands.Register() method. 
Example:
```go
commands.Register("login", cmds.HandlerLogin)
```
Logged-in command handlers need to be wrapped in the middleware function to be passed into the Register method.
Example:
```go
commands.Register("addfeed", cmds.MiddlewareLoggedIn(cmds.HandlerAddFeed))
```

Once the user input in parsed in main.go the commands.Run() method uses the commands map to call the appropriate command using the registered handlers.

# Useful links
- [What is an RSS Feed?](https://en.wikipedia.org/wiki/RSS) 
- [Go documentation](https://go.dev/)
- [Postgres documentation](https://www.postgresql.org/)
- [SQLC Documentation](https://docs.sqlc.dev/en/latest/index.html)
- [PQ documentation](https://pkg.go.dev/github.com/lib/pq)
