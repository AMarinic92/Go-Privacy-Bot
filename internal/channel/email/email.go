// Package email is a Channel that sends letters to a broker's privacy
// contact over SMTP, from that broker's alias.
package email

import (
	"context"

	"github.com/AMarinic92/Go-Privacy-Bot/internal/engine"
	"github.com/AMarinic92/Go-Privacy-Bot/internal/lifecycle"
)

// Sender sends requests by email.
type Sender struct {
	// TODO: fields
}

var _ engine.Channel = (*Sender)(nil)

// Submit sends request r by email.
func (s *Sender) Submit(ctx context.Context, r lifecycle.Request) (engine.Receipt, error) {
	panic("not implemented")
}
