package tui

import (
	"context"
	"strings"

	"github.com/Nishivaly/go-reddit/v2/reddit"
	"github.com/Nishivaly/terminal-reddit-mod/auth"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	client *reddit.Client
	user   userData
	err    error
	list   list.Model
}

func (m Model) Init() tea.Cmd {
	return initRedditClient
}

type redditClientMsg struct {
	client *reddit.Client
}

func initRedditClient() tea.Msg {
	client, err := auth.GetRedditClient()
	if err != nil {
		return errMsg{err}
	}
	return redditClientMsg{client}
}

type userData struct {
	moderated []*reddit.Subreddit
	modQueue  modQueue
}

func (u userData) stringifyModerated() string {
	names := make([]string, len(u.moderated))
	for i, sub := range u.moderated {
		names[i] = sub.Name
	}
	return strings.Join(names, "+")
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

type redditModeratedMsg struct {
	moderated []*reddit.Subreddit
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }
