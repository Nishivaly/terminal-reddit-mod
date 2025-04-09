package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case redditClientMsg:
		m.client = msg.client
		return m, getModerated(m.client)

	case redditModeratedMsg:
		m.user.moderated = msg.moderated
		return m, getModQueue(m.client, m.user.stringifyModerated())

	case redditModQueueMsg:
		m.user.modQueue.posts = msg.posts
		m.user.modQueue.comments = msg.comments
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
