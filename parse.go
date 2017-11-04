package metro2

import (
	"fmt"
)

// The Metro2 interface defines types that can format themselves according to the Metro2 format specification.
type Metro2 interface {
	Metro(int) string
}

// Identify the type of record and return the appropriate data structure.
func parseFixed(source string) (Metro2, error) {
	e := errParser{source: source}

	rdw := e.parseNumber(1, 4)
	identifier := e.parseText(5, 10)

	if rdw != len(source) {
		return nil, fmt.Errorf("Reported record length: (%d) does not match actual length of record: (%d)", rdw, len(source))
	}

	if identifier == "HEADER" {
		return parseFixedHeader(source)
	}

	return parseFixedBase(source)
}
