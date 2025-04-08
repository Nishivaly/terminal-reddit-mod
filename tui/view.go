package tui

import (
	"fmt"
)

func (m Model) View() string {
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

	if m.user != nil {
		for _, subreddit := range m.user.moderated {
			s += fmt.Sprintf("%s\n", subreddit.Name)
		}
	}
	return "\n" + s + "\n\n"
}
