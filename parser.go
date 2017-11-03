package metro2

import (
	"errors"
	"fmt"
	"strconv"
)

type Header struct {
}

func parseFixedHeader(source string) (*Header, error) {
	header := &Header{}
	e := errParser{}

	rdw := e.parseNumeric(source[0:4])
	if rdw != len(source) {
		return nil, errors.New(fmt.Sprintf("Reported record length: (%d) does not match actual length of record: (%d).", rdw, len(source)))
	}

	return header, e.err
}

type errParser struct {
	err error
}

func (e *errParser) parseNumeric(str string) (n int) {
	if e.err != nil {
		return n
	}

	n, err := strconv.Atoi(str)
	if err != nil {
		e.err = err
	}
	return n
}
