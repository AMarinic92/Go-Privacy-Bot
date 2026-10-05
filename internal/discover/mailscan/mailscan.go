// Package mailscan is a Source that reads the operator's mailboxes to list
// companies that already hold the operator's data. A later mass-unsubscribe
// feature would reuse it.
package mailscan

import (
	"context"

	"github.com/AMarinic92/Go-Privacy-Bot/internal/broker"
	"github.com/AMarinic92/Go-Privacy-Bot/internal/engine"
	"github.com/AMarinic92/Go-Privacy-Bot/internal/profile"
)

// Scanner reads mailboxes for companies that hold the operator's data.
type Scanner struct {
	// TODO: fields
}

var _ engine.Source = (*Scanner)(nil)

// Find reports what the operator's mailboxes show about broker b.
func (s *Scanner) Find(ctx context.Context, p profile.Profile, b broker.Broker) ([]engine.Listing, error) {
	panic("not implemented")
}
