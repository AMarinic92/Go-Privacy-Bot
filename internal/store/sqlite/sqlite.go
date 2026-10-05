// Package sqlite is a Store backed by a local SQLite database. It keeps
// every request, reply and timestamp. That log is the evidence file for a
// complaint.
package sqlite

import (
	"context"
	"time"

	"github.com/AMarinic92/Go-Privacy-Bot/internal/engine"
	"github.com/AMarinic92/Go-Privacy-Bot/internal/lifecycle"
)

// Store keeps request records and events in SQLite.
type Store struct {
	// TODO: fields
}

var _ engine.Store = (*Store)(nil)

// Due returns the request records that are due at time now.
func (s *Store) Due(ctx context.Context, now time.Time) ([]lifecycle.Request, error) {
	panic("not implemented")
}

// Save stores request record r with event e.
func (s *Store) Save(ctx context.Context, r lifecycle.Request, e lifecycle.Event) error {
	panic("not implemented")
}
