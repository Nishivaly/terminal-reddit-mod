package tui

import (
	"context"
	"fmt"

	"github.com/Nishivaly/go-reddit/v2/reddit"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type modQueue struct {
	posts    []*reddit.Post
	comments []*reddit.Comment
}

type modQueueItem struct {
	title string
	body  string
	id    string // For future reference (e.g., approving/removing)
	kind  string // "post" or "comment"
}

func (m modQueueItem) Title() string       { return m.title }
func (m modQueueItem) Description() string { return m.body }
func (m modQueueItem) FilterValue() string { return m.title }

type redditModQueueMsg struct {
	posts    []*reddit.Post
	comments []*reddit.Comment
}

type modQueueListMsg struct {
	modQueue []list.Item
}

func getModQueue(client *reddit.Client, subreddits string) tea.Cmd {
	return func() tea.Msg {
		posts, comments, _, err := client.Moderation.Queue(context.Background(), subreddits, &reddit.ListOptions{})
		if err != nil {
			return errMsg{err}
		}
		return redditModQueueMsg{posts, comments}
	}
}

func buildModQueueList(modQueue modQueue) tea.Cmd {
	return func() tea.Msg {
		var modQueueList []list.Item
		for _, comment := range modQueue.comments {
			modQueueList = append(modQueueList, modQueueItem{
				title: fmt.Sprintf("Comment by %s", comment.Author),
				body:  comment.Body,
				id:    comment.ID,
				kind:  "comment",
			})
		}
		for _, post := range modQueue.posts {
			modQueueList = append(modQueueList, modQueueItem{
				title: fmt.Sprintf("Post: %s", post.Title),
				body:  post.Body,
				id:    post.ID,
				kind:  "post",
			})
		}
		return modQueueListMsg{modQueueList}
	}
}
