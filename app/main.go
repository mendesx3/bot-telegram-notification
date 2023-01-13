package main

import (
	"context"
	"fmt"
	twitterscraper "github.com/n0madic/twitter-scraper"
	"time"
)

func main() {
	var seen = make(map[string]bool)

	interval := 5 * time.Second
	query := "#crypto"
	scraper := twitterscraper.New()

	for range time.Tick(interval) {
		for tweet := range scraper.SearchTweets(context.Background(), query, 100) {
			if tweet.Error != nil {
				panic(tweet.Error)
			}
			if seen[tweet.ID] {
				continue
			}
			seen[tweet.ID] = true
			fmt.Println(tweet.Text)
		}
	}
}
