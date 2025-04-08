package tui

import (
	"github.com/Nishivaly/go-reddit/v2/reddit"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	client *reddit.Client
  user *userData
	posts  []*reddit.Post
	err    error
}

type userData struct {
  moderated []*reddit.Subreddit
}

func (m Model) Init() tea.Cmd {
	return initRedditClient
}
