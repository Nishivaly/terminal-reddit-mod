package tui

import (
	"context"

	"github.com/Nishivaly/go-reddit/v2/reddit"
	"github.com/Nishivaly/terminal-reddit-mod/auth"
	tea "github.com/charmbracelet/bubbletea"
)

func initRedditClient() tea.Msg {
	client, err := auth.GetRedditClient()
	if err != nil {
		return errMsg{err}
	}
	return redditClientMsg{client}
}

func getPosts(client *reddit.Client) tea.Cmd {
	return func() tea.Msg {
		posts, _, err := client.Subreddit.TopPosts(context.Background(), "golang", &reddit.ListPostOptions{
			ListOptions: reddit.ListOptions{
				Limit: 5,
			},
			Time: "all",
		})
		if err != nil {
			return errMsg{err}
		}
		return redditPostsMsg{posts}
	}
}
