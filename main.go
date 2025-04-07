package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Nishivaly/go-reddit/v2/reddit"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	client *reddit.Client
	posts  []*reddit.Post
	err    error
}

type redditClientMsg struct {
	client *reddit.Client
}

type redditPostsMsg struct {
	posts []*reddit.Post
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func initRedditClient() tea.Msg {
	client, err := getRedditClient()
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

func (m model) Init() tea.Cmd {
	return initRedditClient
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nWe had some trouble: %s\n\n", m.err.Error())
	}

	s := fmt.Sprintf("Connecting bruh")

	if m.client != nil {
		s += fmt.Sprint("\nconnecto bruh")
	}

	if m.posts != nil {
		s += "\nPosts:\n"
		for i, post := range m.posts {
			s += fmt.Sprintf("%d. %s\n", i+1, post.Title)
		}
	}
	return "\n" + s + "\n\n"
}

func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
