package cmds

import (
	"fmt"
	"log"
	"time"
	"strings"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"bootdev/gator/internal/rss"
	"bootdev/gator/internal/database"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Expected 1 argument: time between requests")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Invalid duration provided as argument. Err: %w", err)
	}
	fmt.Printf("Collecting feeds every %s\n", cmd.Args[0])

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func scrapeFeeds(s *State) error {
	nextFeedToFetch, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't get next feed to fetch. Err: %w", err)
	}

	params := database.MarkFeedFetchedParams{
		LastFetchedAt:	sql.NullTime{Time: time.Now(), Valid: true},
		ID:				nextFeedToFetch.ID,
	}
	err = s.Db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Couldn't mark feed as fetched. Err: %w", err)
	}

	feed, err := rss.FetchFeed(context.Background(), nextFeedToFetch.Url)
	if err != nil {
		return fmt.Errorf("Couldn't collect feed %s. Err: %w", nextFeedToFetch.Url, err)
	}
	for _, item := range feed.Channel.Item {
		hasTitle := hasField(item.Title)
		if !hasTitle {
			continue
		}
		hasLink := hasField(item.Link)
		if !hasLink {
			continue
		}
		hasDescription := hasField(item.Description)
		hasPubDate := hasField(item.PubDate)
		parsedPubDate := time.Time{}

		if hasPubDate {
			parsedPubDate, err = parsePubDate(item.PubDate)
			if err != nil {
				hasPubDate = false
				err = fmt.Errorf("Failed to parse pubDate %q from feed %q: %w", item.PubDate, nextFeedToFetch.Url, err)
				log.Println(err)
			}
		}
		
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: hasDescription},
			PublishedAt: sql.NullTime{Time: parsedPubDate, Valid: hasPubDate},
			FeedID:      nextFeedToFetch.ID,
		}

		_, err := s.Db.CreatePost(context.Background(), params)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post. Err: %v", err)
		}
	}
	return nil
}

func hasField(field string) bool {
	if field != ""{
		return true
	}
	return false
}

func parsePubDate(pubDate string) (time.Time, error) {
	layouts := []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC822,
		time.RFC850,
		time.RFC822Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.UnixDate,
		time.ANSIC,
		time.RubyDate,
		time.Layout,
		time.DateTime,
		time.DateOnly,
	}

	parsedPubDate := time.Time{}
	var err error
	for _, layout := range layouts {
		parsedPubDate, err = time.Parse(layout, pubDate)
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		return time.Time{}, err
	}

	return parsedPubDate, nil
}