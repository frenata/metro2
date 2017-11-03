package metro2

import (
	"fmt"
	"time"
)

type Header struct {
	Cycle        string    `json:"cycle"`
	ActivityDate time.Time `json:"activity_date"`
	DateCreated  time.Time `json:"created_date"`
	Agencies     Agencies  `json:"agency_numbers"`
	Program      Program   `json:"program"`
	Reporter     Reporter  `json:"reporter"`
	Software     Software  `json:"software"`
}

func (h Header) metro2() {}

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
	Name            string `json:"name"`
	Address         string `json:"address"`
	TelephoneNumber int    `json:"telephone_number"`
}

type Software struct {
	VendorName string `json:"vendor_name"`
	Version    string `json:"version"`
}

func parseFixedHeader(source string) (*Header, error) {
	e := errParser{source: source}

	cycle := e.parseText(11, 12)
	activity := e.parseDate(date, 48, 55)
	created := e.parseDate(date, 56, 63)

	agencies := Agencies{e.parseText(13, 22), e.parseText(23, 32), e.parseText(33, 37), e.parseText(38, 47)}

	program := Program{e.parseDate(date, 64, 71), e.parseDate(date, 72, 79)}

	reporter := Reporter{e.parseText(80, 119), e.parseText(120, 215), e.parseNumber(216, 225)}

	software := Software{e.parseText(226, 265), e.parseText(266, 270)}

	if e.err != nil {
		return nil, e.err
	}

	return &Header{cycle, activity, created, agencies, program, reporter, software}, nil
}

func formatFixedHeader(h *Header, length int) string {
	prefix := fmt.Sprintf("%04dHEADER", length)

	agencies := fmt.Sprintf("%-10s%-10s%-5s%-10s", h.Agencies.Innovis, h.Agencies.Equifax, h.Agencies.Experian, h.Agencies.TransUnion)
	program := fmt.Sprintf("%s%s", h.Program.StartDate.Format(date), h.Program.RevisionDate.Format(date))
	reporter := fmt.Sprintf("%-40s%-96s%10d", h.Reporter.Name, h.Reporter.Address, h.Reporter.TelephoneNumber)
	software := fmt.Sprintf("%-40s%-5s", h.Software.VendorName, h.Software.Version)

	reserved := fmt.Sprintf("%-156s", "")

	header := fmt.Sprintf("%-2s%s%s%s%s%s%s%s\n", h.Cycle, agencies, h.ActivityDate.Format(date), h.DateCreated.Format(date), program, reporter, software, reserved)

	return prefix + header
}
