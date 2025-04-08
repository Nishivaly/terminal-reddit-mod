package tui

import (
	"github.com/Nishivaly/go-reddit/v2/reddit"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	client *reddit.Client
	posts  []*reddit.Post
	err    error
}

func (m Model) Init() tea.Cmd {
	return initRedditClient
}
