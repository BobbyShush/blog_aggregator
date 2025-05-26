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
```
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
`go install gator`

- Install Goose:
`go install github.com/pressly/goose/v3/cmd/goose@latest`
- Run up the migration. Use this command from the base of the repository with the proper credentials
MacOS - System username, no password
Linux/WSL - Username "postgres", in-app password (set earlier)
`goose -dir ./sql/schema postgres postgres://<username>:<password>@localhost:5432/gator up`

### Config setup
In the home directory, create a file named .gatorconfig.json and paste this content inside (adapt the username and password accordingly):
```
{"db_url":"postgres://<username>:<password>@localhost:5432/gator?sslmode=disable","current_user_name":""}
```

## Command list
### Register

### Login

### Users

### Reset (for testing only)

### Add feed

### Feeds

### Follow

### Following

### Unfollow

### Agg

### Browse

# Dev guide
## Commands and how they are implemented
General command implementation (and how to add more if needed)

## Errors
Error flow and handling (so that I can improve if I stumble on new ones)

## Useful links
- What is an RSS Feed?
- Go
- Postgres
- Sqlc
- Pq
