// Package engine runs the daily tick and picks the next action for each
// broker.
//
// The engine holds every decision. It reaches the outside world only through
// the interfaces in ports.go and imports no adapter package.
package engine

import (
	"context"
	"time"
)

// Engine runs the daily tick. It will hold the five ports: Source, Channel,
// Inbox, Store and Notifier.
type Engine struct {
	// TODO: fields
}

// Tick runs one daily pass at time now. It moves each request record at most
// one step.
func (e *Engine) Tick(ctx context.Context, now time.Time) error {
	panic("not implemented")
}
