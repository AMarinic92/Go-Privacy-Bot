// Package broker defines the broker definition schema and its validation.
//
// A definition describes one data broker or people-search site. Definitions
// are YAML files under brokers/ in the repository. The format is a draft; see
// brokers/README.md.
package broker

// Broker is one parsed broker definition.
type Broker struct {
	// TODO: fields
}

// Parse decodes one broker definition from data. It takes bytes, not a file
// path.
func Parse(data []byte) (Broker, error) {
	panic("not implemented")
}

// Validate reports whether b is a complete and consistent definition.
func (b Broker) Validate() error {
	panic("not implemented")
}
