package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Nishivaly/go-reddit/v2/reddit"
)

func main() {
	client, err := getRedditClient()
	if err != nil {
		panic(err)
	}

	posts, _, err := client.Subreddit.TopPosts(context.Background(), "golang", &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 5,
		},
		Time: "all",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received %d posts.\n", len(posts))
}
