package metro2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Mon Jan 2 15:04:05 -0700 MST 2006
const dateFormat string = "01022006"

type Header struct {
	Cycle        int       `json:"cycle"`
	ActivityDate time.Time `json:"activity_date"`
	DateCreated  time.Time `json:"created_date"`
	Agencies     Agencies  `json:"agency_numbers"`
	Program      Program   `json:"program"`
	Reporter     Reporter  `json:"reporter"`
	Software     Software  `json:"software"`
}

type Agencies struct {
	Innovis    string `json:"innovis_number"`
	Equifax    string `json:"equifax_number"`
	Experian   string `json:"experian_number"`
	TransUnion string `json:"transunion_number"`
}

type Program struct {
	StartDate    time.Time `json:"start_date"`
	RevisionDate time.Time `json:"revision_date"`
}

type Reporter struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber int    `json:phone_number"`
}

type Software struct {
	VendorName string `json:"vendor_name"`
	Version    string `json:"version"`
}

func parseFixedHeader(source string) (*Header, error) {
	e := errParser{source: source}

	rdw := e.parseNumeric(1, 4)
	if rdw != len(source) {
		return nil, errors.New(fmt.Sprintf("Reported record length: (%d) does not match actual length of record: (%d).", rdw, len(source)))
	}

	cycle := e.parseNumeric(11, 12)
	activity := e.parseDate(48, 55)
	created := e.parseDate(56, 63)

	agencies := Agencies{e.parseText(13, 22), e.parseText(23, 32), e.parseText(33, 37), e.parseText(38, 47)}

	program := Program{e.parseDate(64, 71), e.parseDate(72, 79)}

	reporter := Reporter{e.parseText(80, 119), e.parseText(120, 215), e.parseNumeric(216, 225)}

	software := Software{e.parseText(226, 265), e.parseText(266, 270)}

	if e.err != nil {
		return nil, e.err
	}

	return &Header{cycle, activity, created, agencies, program, reporter, software}, nil
}

type errParser struct {
	source string
	err    error
}

func (e *errParser) parseNumeric(start, end int) (n int) {
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
