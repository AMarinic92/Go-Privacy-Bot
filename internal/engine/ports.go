package engine

import (
	"context"
	"time"

	"github.com/AMarinic92/Go-Privacy-Bot/internal/broker"
	"github.com/AMarinic92/Go-Privacy-Bot/internal/lifecycle"
	"github.com/AMarinic92/Go-Privacy-Bot/internal/profile"
)

// Source finds the operator's listings on one broker.
type Source interface {
	Find(ctx context.Context, p profile.Profile, b broker.Broker) ([]Listing, error)
}

// Channel submits a request to a broker.
type Channel interface {
	Submit(ctx context.Context, r lifecycle.Request) (Receipt, error)
}

// Inbox fetches the messages that arrived since a given time.
type Inbox interface {
	Fetch(ctx context.Context, since time.Time) ([]Message, error)
}

// Store loads the request records that are due and saves each record with
// the event that changed it.
type Store interface {
	Due(ctx context.Context, now time.Time) ([]lifecycle.Request, error)
	Save(ctx context.Context, r lifecycle.Request, e lifecycle.Event) error
}

// Notifier sends the operator a digest.
type Notifier interface {
	Send(ctx context.Context, d Digest) error
}

// Listing is a match that a Source found on a broker.
type Listing struct {
	// TODO: fields
}

// Receipt is what a Channel returns after it submits a request.
type Receipt struct {
	// TODO: fields
}

// Message is one message fetched from an Inbox, such as a broker reply or a
// confirmation link.
type Message struct {
	// TODO: fields
}

// Digest is the report sent to the operator, including the manual queue.
type Digest struct {
	// TODO: fields
}
