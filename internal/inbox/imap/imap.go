// Package imap is an Inbox that polls the catch-all IMAP mailbox. It finds
// confirmation links and classifies broker replies.
package imap

import (
	"context"
	"time"

	"github.com/AMarinic92/Go-Privacy-Bot/internal/engine"
)

// Client reads the catch-all mailbox.
type Client struct {
	// TODO: fields
}

var _ engine.Inbox = (*Client)(nil)

// Fetch returns the messages that arrived since the given time.
func (c *Client) Fetch(ctx context.Context, since time.Time) ([]engine.Message, error) {
	panic("not implemented")
}
