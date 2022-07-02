package main

import (
	"fmt"
	"news-feeder/platform/newsfeed"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", handler.PingGet())
	// r.Run()

	feed := newsfeed.New()

	fmt.Println(feed)

	feed.Add(newsfeed.Item{"Hello", "How r u doing"})

	fmt.Println(feed)
}
