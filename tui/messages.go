package tui

import (
	"github.com/Nishivaly/go-reddit/v2/reddit"
)

type redditClientMsg struct {
	client *reddit.Client
}

type redditModeratedMsg struct {
	moderated []*reddit.Subreddit
}

type redditModQueueMsg struct {
	posts    []*reddit.Post
	comments []*reddit.Comment
}

type redditPostsMsg struct {
	posts []*reddit.Post
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }
