// Package lifecycle defines the states of a request record and the rule that
// moves a record from one state to the next.
//
// Each broker has one request record. A daily run moves a record at most one
// step. The state diagram is in docs/architecture.md.
package lifecycle

import "time"

// State is the position of a request record in its lifecycle.
type State string

// States, in the order they appear in the state diagram.
const (
	Candidate        State = "Candidate"        // broker definition loaded
	NotFound         State = "NotFound"         // no listing found; re-checked when due
	Listed           State = "Listed"           // a Source found a listing
	ManualQueue      State = "ManualQueue"      // needs ID, phone or CAPTCHA from the operator
	Sent             State = "Sent"             // submitted by email or form, or marked done by the operator
	AwaitingConfirm  State = "AwaitingConfirm"  // the broker emailed a confirmation link
	AwaitingResponse State = "AwaitingResponse" // waiting on the 30-day response clock
	Verify           State = "Verify"           // the reply says removed; checking the listing
	Overdue          State = "Overdue"          // 30 days with no reply, or still listed
	Removed          State = "Removed"          // listing gone
	Escalated        State = "Escalated"        // reminder ignored; the bot stops here
)

// Request is the request record for one broker.
type Request struct {
	// TODO: fields
}

// Event is something that happened to a request record. Every event is
// stored as evidence.
type Event struct {
	// TODO: fields
}

// Next returns the state that r moves to when e happens at time now. It
// moves r at most one step.
func Next(r Request, e Event, now time.Time) (State, error) {
	panic("not implemented")
}
