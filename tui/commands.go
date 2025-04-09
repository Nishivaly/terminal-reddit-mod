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

// func getPosts(client *reddit.Client) tea.Cmd {
// 	return func() tea.Msg {
// 		posts, _, err := client.Subreddit.TopPosts(context.Background(), "golang", &reddit.ListPostOptions{
// 			ListOptions: reddit.ListOptions{
// 				Limit: 5,
// 			},
// 			Time: "all",
// 		})
// 		if err != nil {
// 			return errMsg{err}
// 		}
// 		return redditPostsMsg{posts}
// 	}
// }

func getModQueue(client *reddit.Client, subreddits string) tea.Cmd {
	return func() tea.Msg {
		posts, comments, _, err := client.Moderation.Queue(context.Background(), subreddits, &reddit.ListOptions{})
		if err != nil {
			return errMsg{err}
		}
		return redditModQueueMsg{posts, comments}
	}
}

func getModerated(client *reddit.Client) tea.Cmd {
	return func() tea.Msg {
		moderated, _, err := client.Subreddit.Moderated(context.Background(), &reddit.ListSubredditOptions{})
		if err != nil {
			return errMsg{err}
		}
		return redditModeratedMsg{moderated}
	}
}
