// Package registry turns public broker lists into draft broker definitions.
// Every draft is reviewed by hand before it becomes a definition.
package registry

import (
	"io"

	"github.com/AMarinic92/Go-Privacy-Bot/internal/broker"
)

// ImportCSV reads a broker list in CSV form, such as the California data
// broker registry CSV, and returns draft definitions.
func ImportCSV(r io.Reader) ([]broker.Broker, error) {
	panic("not implemented")
}
