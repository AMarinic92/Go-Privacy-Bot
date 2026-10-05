// Package notify is a Notifier that sends the operator the digest by push or
// email.
package notify

import (
	"context"

	"github.com/AMarinic92/Go-Privacy-Bot/internal/engine"
)

// Notifier sends digests to the operator.
type Notifier struct {
	// TODO: fields
}

var _ engine.Notifier = (*Notifier)(nil)

// Send delivers digest d to the operator.
func (n *Notifier) Send(ctx context.Context, d engine.Digest) error {
	panic("not implemented")
}
