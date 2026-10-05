// Package manual is a Channel that queues a task for the operator instead of
// contacting the broker. It covers ID uploads, phone calls and CAPTCHAs.
package manual

import (
	"context"

	"github.com/AMarinic92/Go-Privacy-Bot/internal/engine"
	"github.com/AMarinic92/Go-Privacy-Bot/internal/lifecycle"
)

// Queue holds tasks for the operator.
type Queue struct {
	// TODO: fields
}

var _ engine.Channel = (*Queue)(nil)

// Submit queues request r as a task for the operator.
func (q *Queue) Submit(ctx context.Context, r lifecycle.Request) (engine.Receipt, error) {
	panic("not implemented")
}
