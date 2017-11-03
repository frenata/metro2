package metro2

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

//Go Date String: Mon Jan 2 15:04:05 -0700 MST 2006
const dateFormat string = "01022006" // aka MMDDYYYY

// errParser allows for clean parsing of a long string into values
// by retaining the first error encountered
type errParser struct {
	source string
	err    error
}

// attempt to parse a string into a number
func (e *errParser) parseNumber(start, end int) (n int) {
	if e.err != nil {
		return n
	}

	str := e.subStr(start, end)
	n, err := strconv.Atoi(str)
	if err != nil {
		e.err = err
	}
	return n
}

// attempt to parse a string into a date
func (e *errParser) parseDate(start, end int) (date time.Time) {
	if e.err != nil {
		return date
	}

	str := e.subStr(start, end)
	date, err := time.Parse(dateFormat, str)
	if err != nil {
		e.err = err
	}
	return date
}

// parse a string into a string by removing space padding
func (e *errParser) parseText(start, end int) (text string) {
	if e.err != nil {
		return text
	}

	str := e.subStr(start, end)
	text = strings.TrimSpace(str)

	return text
}

// get positions n - m, inclusive, 1 indexed
func (e errParser) subStr(n, m int) string {
	if m > len(e.source) {
		e.err = errors.New("Tried to index a source file beyond length" + strconv.Itoa(m))
		return ""
	}

	return e.source[n-1 : m]
}
