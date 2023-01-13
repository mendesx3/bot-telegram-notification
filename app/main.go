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
	clearInterval := 2 * time.Hour
	query := "#crypto"
	scraper := twitterscraper.New()

	ticker := time.NewTicker(interval)
	clearTicker := time.NewTicker(clearInterval)

	go func() {
		for {
			select {
			case <-ticker.C:
				for tweet := range scraper.SearchTweets(context.Background(), query, 50) {
					if tweet.Error != nil {
						fmt.Println(tweet.Error)
						continue
					}
					if seen[tweet.ID] {
						continue
					}
					seen[tweet.ID] = true
					fmt.Println(tweet.Text)
				}
			case <-clearTicker.C:
				fmt.Println("clearing seen map")
				seen = make(map[string]bool)
			}
		}
	}()

	<-time.After(10 * time.Minute)
	ticker.Stop()
	clearTicker.Stop()
}
