package tui

import (
	"fmt"
)

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nWe had some trouble: %s\n\n", m.err.Error())
	}

	if m.client == nil {
		return fmt.Sprintf("Logging in to Reddit...")
	}

	s := fmt.Sprint("\nLogged in.\n")

	s += "\nModerated subs:\n"
	for _, subreddit := range m.user.moderated {
		s += fmt.Sprintf("%s\n", subreddit.Name)
	}

	if len(m.user.modQueue.posts) > 0 {
		s += "\nModqueue posts:\n"
		for i, post := range m.user.modQueue.posts {
			s += fmt.Sprintf("%d. %s\n", i+1, post.Title)
		}
	}

	if len(m.user.modQueue.comments) > 0 {
		s += "\nModqueue comments:\n"
		for i, comment := range m.user.modQueue.comments {
			s += fmt.Sprintf("%d. Author: %s - Body: %s\n", i+1, comment.Author, comment.Body)
		}
	}

	return s + "\n\n"
}
