package metro2

import (
	"fmt"
	"time"
)

// A Header contains data related to the sender of the Base segments to follow.
type Header struct {
	Cycle        string    `json:"cycle"`
	ActivityDate time.Time `json:"activity_date"`
	DateCreated  time.Time `json:"created_date"`
	Agencies     agencies  `json:"agency_numbers"`
	Program      program   `json:"program"`
	Reporter     reporter  `json:"reporter"`
	Software     software  `json:"software"`
}

// Metro creates a string in the Metro2 data from the Header data structure.
func (h Header) Metro(length int) string {
	prefix := fmt.Sprintf("%04dHEADER", length)

	reserved := fmt.Sprintf("%-156s", "")

	header := fmt.Sprintf("%-2s%s%s%s%s%s%s%s\n", h.Cycle, h.Agencies.metro(), h.ActivityDate.Format(date), h.DateCreated.Format(date), h.Program.metro(), h.Reporter.metro(), h.Software.metro(), reserved)

	return prefix + header
}

type agencies struct {
	Innovis    string `json:"innovis_number"`
	Equifax    string `json:"equifax_number"`
	Experian   string `json:"experian_number"`
	TransUnion string `json:"transunion_number"`
}

func (a agencies) metro() string {
	return fmt.Sprintf("%-10s%-10s%-5s%-10s", a.Innovis, a.Equifax, a.Experian, a.TransUnion)
}

type program struct {
	StartDate    time.Time `json:"start_date"`
	RevisionDate time.Time `json:"revision_date"`
}

func (p program) metro() string {
	return fmt.Sprintf("%s%s", p.StartDate.Format(date), p.RevisionDate.Format(date))
}

type reporter struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	TelephoneNumber int    `json:"telephone_number"`
}

func (r reporter) metro() string {
	return fmt.Sprintf("%-40s%-96s%10d", r.Name, r.Address, r.TelephoneNumber)
}

type software struct {
	VendorName string `json:"vendor_name"`
	Version    string `json:"version"`
}

func (s software) metro() string {
	return fmt.Sprintf("%-40s%-5s", s.VendorName, s.Version)
}

// Parse a fixed length header record.
func parseFixedHeader(source string) (*Header, error) {
	e := errParser{source: source}

	cycle := e.parseText(11, 12)
	activity := e.parseDate(date, 48, 55)
	created := e.parseDate(date, 56, 63)

	agencies := agencies{
		Innovis:    e.parseText(13, 22),
		Equifax:    e.parseText(23, 32),
		Experian:   e.parseText(33, 37),
		TransUnion: e.parseText(38, 47),
	}

	program := program{
		StartDate:    e.parseDate(date, 64, 71),
		RevisionDate: e.parseDate(date, 72, 79),
	}

	reporter := reporter{
		Name:            e.parseText(80, 119),
		Address:         e.parseText(120, 215),
		TelephoneNumber: e.parseNumber(216, 225),
	}

	software := software{
		VendorName: e.parseText(226, 265),
		Version:    e.parseText(266, 270),
	}

	if e.err != nil {
		return nil, e.err
	}

	return &Header{cycle, activity, created, agencies, program, reporter, software}, nil
}
