package tui

import (
	"strings"

	"github.com/Nishivaly/go-reddit/v2/reddit"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	client *reddit.Client
	user   userData
	err    error
}

func (m Model) Init() tea.Cmd {
	return initRedditClient
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

type modQueue struct {
	posts    []*reddit.Post
	comments []*reddit.Comment
}
