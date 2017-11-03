package metro2

import (
	"errors"
	"fmt"
)

// Trivially define a Metro2 interface type to allow for multiple structs to be returned from `parseFixed`
type Metro2 interface {
	metro2()
}

// Identify the type of record and return the appropriate data structure.
func parseFixed(source string) (Metro2, error) {
	e := errParser{source: source}

	rdw := e.parseNumber(1, 4)
	identifier := e.parseText(5, 10)

	if rdw != len(source) {
		return nil, errors.New(fmt.Sprintf("Reported record length: (%d) does not match actual length of record: (%d).", rdw, len(source)))
	}

	if identifier == "HEADER" {
		return parseFixedHeader(source)
	} else {
		return parseFixedBase(source)
	}
	return nil, nil
}
