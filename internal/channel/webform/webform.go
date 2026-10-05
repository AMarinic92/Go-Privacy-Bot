// Package webform is a Channel that submits a request by running a broker
// definition's form steps in headless Chromium.
//
// ID uploads, phone calls and CAPTCHAs go to the manual queue. The runner
// never tries to solve or bypass them.
package webform

import (
	"context"

	"github.com/AMarinic92/Go-Privacy-Bot/internal/engine"
	"github.com/AMarinic92/Go-Privacy-Bot/internal/lifecycle"
)

// Runner runs form steps in headless Chromium.
type Runner struct {
	// TODO: fields
}

var _ engine.Channel = (*Runner)(nil)

// Submit submits request r through the broker's web form.
func (w *Runner) Submit(ctx context.Context, r lifecycle.Request) (engine.Receipt, error) {
	panic("not implemented")
}
