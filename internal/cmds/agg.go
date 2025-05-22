package cmds

import (
	"fmt"
	"context"
	"bootdev/blog_aggregator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("Title: %s\n", feed.Channel.Title)
	fmt.Printf("Link: %s\n", feed.Channel.Link)
	fmt.Printf("Description: %s\n", feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		fmt.Printf("Item #%d\n", i+1)
		fmt.Printf("--- Title: %s\n", item.Title)
		fmt.Printf("--- Link: %s\n", item.Link)
		fmt.Printf("--- Description: %s\n", item.Description)
		fmt.Printf("--- Published: %s\n", item.PubDate)
	}
	return nil
}