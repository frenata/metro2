package metro2

import (
	"encoding/json"
	"errors"
	"time"
)

// Base represents data related to a customer's account, payment history, biographical data, and contact details.
type Base struct {
	Timestamp      time.Time `json:"timestamp"`
	Correction     int       `json:"correction"`
	Identification string    `json:"identification"`
	Cycle          string    `json:"cycle"`
	Account        account   `json:"account"`
	Consumer       consumer  `json:"consumer"`
}

func (b Base) String() string {
	json, _ := json.MarshalIndent(b, "", "  ")
	return string(json)
}

type account struct {
	AccountNumber string     `json:"account_number"`
	PortfolioType string     `json:"portfolio_type"`
	AccountType   string     `json:"account_type"`
	DateOpened    time.Time  `json:"date_opened"`
	CreditLimit   int        `json:"credit_limit"`
	HighestCredit int        `json:"highest_credit"`
	Terms         terms      `json:"terms"`
	Payment       payment    `json:"payment"`
	Compliance    compliance `json:"compliance"`
}

type terms struct {
	Duration  string `json:"duration"`
	Frequency string `json:"frequency"`
}

type payment struct {
	ScheduledMonthlyAmount int    `json:"scheduled_monthly_amount"`
	ActualAmount           int    `json:"actual_amount"`
	Status                 string `json:"status"`
	Rating                 string `json:"rating"`
	HistoryProfile         string `json:"history_profile"`
	SpecialComment         string `json:"special_comment"`
}

type compliance struct {
	ConditionCode           string    `json:"condition_code"`
	CurrentBalance          int       `json:"current_balance"`
	AmountPastDue           int       `json:"amount_past_due"`
	OriginalChargeOffAmount int       `json:"original_charge_off_amount"`
	BillingDate             time.Time `json:"billing_date"`
	DateFirstDelinquency    time.Time `json:"date_first_delinquency"`
	DateClosed              time.Time `json:"date_closed"`
	DateLastPayment         time.Time `json:"date_last_payment"`
}

type consumer struct {
	TransactionType      string    `json:"transaction_type"`
	Surname              string    `json:"surname"`
	FirstName            string    `json:"first_name"`
	MiddleName           string    `json:"middle_name"`
	GenerationCode       string    `json:"generation_code"`
	SocialSecurityNumber int       `json:"social_security_number"`
	DateOfBirth          time.Time `json:"date_of_birth"`
	ECOACode             string    `json:"ecoa_code"`
	InformationIndicator string    `json:"information_indicator"`
	Contact              contact   `json:"contact"`
}

type contact struct {
	TelephoneNumber int     `json:"telephone_number"`
	Country         string  `json:"country"`
	Address         address `json:"address"`
}

type address struct {
	FirstLine     string `json:"first_line"`
	SecondLine    string `json:"second_line"`
	City          string `json:"city"`
	State         string `json:"state"`
	PostalCode    string `json:"postal_code"`
	Indicator     string `json:"indicator"`
	ResidenceCode string `json:"residence_code"`
}

// Further Work: implement Metro formatting for Base Segment
func (b Base) Metro(length int) string {
	return ""
}

func parseFixedBase(source string) (*Base, error) {
	e := errParser{source: source}

	processing := e.parseNumber(5, 5)
	if processing != 1 {
		return nil, errors.New("Processing indicator was not set to 1")
	}

	time := e.parseDate(timestamp, 6, 19)
	correction := e.parseNumber(20, 20)
	identification := e.parseText(21, 40)
	cycle := e.parseText(41, 42)

	terms := terms{
		Duration:  e.parseText(102, 104),
		Frequency: e.parseText(105, 105),
	}

	payment := payment{
		ScheduledMonthlyAmount: e.parseNumber(106, 114),
		ActualAmount:           e.parseNumber(115, 123),
		Status:                 e.parseText(124, 125),
		Rating:                 e.parseText(126, 126),
		HistoryProfile:         e.parseText(127, 150),
		SpecialComment:         e.parseText(151, 152),
	}

	compliance := compliance{
		ConditionCode:           e.parseText(153, 154),
		CurrentBalance:          e.parseNumber(155, 163),
		AmountPastDue:           e.parseNumber(164, 172),
		OriginalChargeOffAmount: e.parseNumber(173, 181),
		BillingDate:             e.parseDate(date, 182, 189),
		DateFirstDelinquency:    e.parseDate(date, 190, 197),
		DateClosed:              e.parseDate(date, 198, 205),
		DateLastPayment:         e.parseDate(date, 206, 213),
	}

	account := account{
		AccountNumber: e.parseText(43, 72),
		PortfolioType: e.parseText(73, 73),
		AccountType:   e.parseText(74, 75),
		DateOpened:    e.parseDate(date, 76, 83),
		CreditLimit:   e.parseNumber(84, 92),
		HighestCredit: e.parseNumber(93, 101),
		Terms:         terms,
		Payment:       payment,
		Compliance:    compliance,
	}

	address := address{
		FirstLine:     e.parseText(330, 361),
		SecondLine:    e.parseText(362, 393),
		City:          e.parseText(394, 413),
		State:         e.parseText(414, 415),
		PostalCode:    e.parseText(416, 424),
		Indicator:     e.parseText(425, 425),
		ResidenceCode: e.parseText(426, 426),
	}

	contact := contact{
		TelephoneNumber: e.parseNumber(315, 324),
		Country:         e.parseText(328, 329),
		Address:         address,
	}

	consumer := consumer{
		TransactionType:      e.parseText(231, 231),
		Surname:              e.parseText(232, 256),
		FirstName:            e.parseText(257, 276),
		MiddleName:           e.parseText(277, 296),
		GenerationCode:       e.parseText(297, 297),
		SocialSecurityNumber: e.parseNumber(298, 306),
		DateOfBirth:          e.parseDate(date, 307, 314),
		ECOACode:             e.parseText(325, 325),
		InformationIndicator: e.parseText(326, 327),
		Contact:              contact,
	}

	if e.err != nil {
		return nil, e.err
	}

	return &Base{time, correction, identification, cycle, account, consumer}, nil
}
