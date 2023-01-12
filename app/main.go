package main

import (
	"context"
	"fmt"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	func1()
	func2()
	func3()
	func4()

}

func func1() {
	scraper := twitterscraper.New()

	for tweet := range scraper.GetTweets(context.Background(), "", 10) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		fmt.Println(tweet.Text)
	}
}

func func2() {
	scraper := twitterscraper.New()
	tweet, err := scraper.GetTweet("")
	if err != nil {
		panic(err)
	}
	fmt.Println(tweet.Text)
}

func func3() {
	scraper := twitterscraper.New()

	for tweet := range scraper.SearchTweets(context.Background(),
		"#crypto", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		fmt.Println(tweet.Text)
	}
}

func func4() {
	scraper := twitterscraper.New()
	trends, err := scraper.GetTrends()
	if err != nil {
		panic(err)
	}
	fmt.Println(trends)
}

//func metodo5() {
//
//	// search tweets with hashtag golang
//	tweets, err := twitter_scraper.ScrapeSearch("#golang", 100)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	for _, tweet := range tweets {
//		fmt.Println(tweet.User.Username, ": ", tweet.Text)
//	}
//}
