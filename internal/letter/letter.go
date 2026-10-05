// Package letter renders the letters the bot sends to brokers.
//
// Templates cite PIPEDA today. They must be swappable, because Bill C-36
// would add a deletion right if it becomes law. Letters never cite CCPA.
package letter

import (
	"github.com/AMarinic92/Go-Privacy-Bot/internal/broker"
	"github.com/AMarinic92/Go-Privacy-Bot/internal/profile"
)

// Kind is the type of letter.
type Kind string

const (
	// Access asks a broker what personal information it holds. A broker
	// with no public search gets one from an alias, with minimal identifiers.
	Access Kind = "Access"

	// Removal asks for access, withdraws consent, and demands destruction
	// of information that is no longer needed.
	Removal Kind = "Removal"

	// Reminder follows up once when 30 days pass without a reply.
	Reminder Kind = "Reminder"
)

// Letter is a rendered letter.
type Letter struct {
	// TODO: fields
}

// Render produces a letter of kind k from profile p to broker b.
func Render(k Kind, p profile.Profile, b broker.Broker) (Letter, error) {
	panic("not implemented")
}
