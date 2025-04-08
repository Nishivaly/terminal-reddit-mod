package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case redditClientMsg:
		m.client = msg.client
		return m, getPosts(m.client)

	case redditPostsMsg:
		m.posts = msg.posts
		return m, nil

	case errMsg:
		m.err = msg.err

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	return m, nil
}
