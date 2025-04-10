package tui

import (
	"github.com/charmbracelet/bubbles/list"
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
		return m, buildModQueueList(m.user.modQueue)

	case modQueueListMsg:
		delegate := list.NewDefaultDelegate()
		newList := list.New(msg.modQueue, delegate, 0, 0)
		newList.Title = "Modqueue"
    newList.SetSize(80, 20)
    m.list = newList
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
