// Package probe is a Source that searches broker sites for the operator's
// listings.
package probe

import (
	"context"

	"github.com/AMarinic92/Go-Privacy-Bot/internal/broker"
	"github.com/AMarinic92/Go-Privacy-Bot/internal/engine"
	"github.com/AMarinic92/Go-Privacy-Bot/internal/profile"
)

// Prober searches broker sites for listings.
type Prober struct {
	// TODO: fields
}

var _ engine.Source = (*Prober)(nil)

// Find searches broker b for listings that match profile p.
func (pr *Prober) Find(ctx context.Context, p profile.Profile, b broker.Broker) ([]engine.Listing, error) {
	panic("not implemented")
}
