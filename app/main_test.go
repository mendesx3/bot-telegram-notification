package main

import (
	"context"
	twitterscraper "github.com/n0madic/twitter-scraper"
	"testing"
	"time"
)

func TestScraper(t *testing.T) {
	var seen = make(map[string]bool)

	interval := 5 * time.Second
	clearInterval := 2 * time.Hour
	query := "#crypto #BTC"
	scraper := twitterscraper.New()

	ticker := time.NewTicker(interval)
	clearTicker := time.NewTicker(clearInterval)

	go func() {
		for {
			select {
			case <-ticker.C:
				for tweet := range scraper.SearchTweets(context.Background(), query, 50) {
					if tweet.Error != nil {
						t.Error(tweet.Error)
						continue
					}
					if seen[tweet.ID] {
						continue
					}
					seen[tweet.ID] = true
				}
			case <-clearTicker.C:
				seen = make(map[string]bool)
			}
		}
	}()

	<-time.After(10 * time.Second)
	ticker.Stop()
	clearTicker.Stop()
}
