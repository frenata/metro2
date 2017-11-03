package metro2

import (
	"errors"
	"time"
)

type Base struct {
	Timestamp      time.Time `json:"timestamp"`
	Correction     int       `json:"correction"`
	Identification string    `json:"identification"`
	Cycle          string    `json:"cycle"`
	Account        Account   `json:"account"`
	Consumer       Consumer  `json:"consumer"`
}

type Account struct {
	AccountNumber string     `json:"account_number"`
	PortfolioType string     `json:"portfolio_type"`
	AccountType   string     `json:"account_type"`
	DateOpened    time.Time  `json:"date_opened"`
	CreditLimit   int        `json:"credit_limit"`
	HighestCredit int        `json:"highest_credit"`
	Terms         Terms      `json:"terms"`
	Payment       Payment    `json:"payment"`
	Compliance    Compliance `json:"compliance"`
}

type Terms struct {
	Duration  string `json:"duration"`
	Frequency string `json:"frequency"`
}

type Payment struct {
	ScheduledMonthlyAmount int    `json:"scheduled_monthly_amount"`
	ActualAmount           int    `json:"actual_amount"`
	Status                 string `json:"status"`
	Rating                 string `json:"rating"`
	HistoryProfile         string `json:"history_profile"`
	SpecialComment         string `json:"special_comment"`
}

type Compliance struct {
	ConditionCode           string    `json:"condition_code"`
	CurrentBalance          int       `json:"current_balance"`
	AmountPastDue           int       `json:"amount_past_due"`
	OriginalChargeOffAmount int       `json:"original_charge_off_amount"`
	BillingDate             time.Time `json:"billing_date"`
	DateFirstDelinquency    time.Time `json:"date_first_delinquency"`
	DateClosed              time.Time `json:"date_closed"`
	DateLastPayment         time.Time `json:"date_last_payment"`
}

type Consumer struct {
	TransactionType      string    `json:"transaction_type"`
	Surname              string    `json:"surname"`
	FirstName            string    `json:"first_name"`
	MiddleName           string    `json:"middle_name"`
	GenerationCode       string    `json:"generation_code"`
	SocialSecurityNumber int       `json:"social_security_number"`
	DateOfBirth          time.Time `json:"date_of_birth"`
	ECOACode             string    `json:"ecoa_code"`
	InformationIndicator string    `json:"information_indicator"`
	Contact              Contact   `json:"contact"`
}

type Contact struct {
	TelephoneNumber int     `json:"telephone_number"`
	Country         string  `json:"country"`
	Address         Address `json:"address"`
}

type Address struct {
	FirstLine     string `json:"first_line"`
	SecondLine    string `json:"second_line"`
	City          string `json:"city"`
	State         string `json:"state"`
	PostalCode    string `json:"postal_code"`
	Indicator     string `json:"indicator"`
	ResidenceCode string `json:"residence_code"`
}

func (b Base) metro2() {}

func parseFixedBase(source string) (*Base, error) {
	e := errParser{source: source}

	processing := e.parseNumber(5, 5)
	if processing != 1 {
		return nil, errors.New("Processing indicator was not set to 1.")
	}

	time := e.parseDate(timestamp, 6, 19)
	correction := e.parseNumber(20, 20)
	identification := e.parseText(21, 40)
	cycle := e.parseText(41, 42)

	terms := Terms{e.parseText(102, 104), e.parseText(105, 105)}
	payment := Payment{e.parseNumber(106, 114), e.parseNumber(115, 123), e.parseText(124, 125), e.parseText(126, 126), e.parseText(127, 150), e.parseText(151, 152)}
	compliance := Compliance{e.parseText(153, 154), e.parseNumber(155, 163), e.parseNumber(164, 172), e.parseNumber(173, 181), e.parseDate(date, 182, 189), e.parseDate(date, 190, 197), e.parseDate(date, 198, 205), e.parseDate(date, 206, 213)}
	account := Account{e.parseText(43, 72), e.parseText(73, 73), e.parseText(74, 75), e.parseDate(date, 76, 83), e.parseNumber(84, 92), e.parseNumber(93, 101), terms, payment, compliance}

	address := Address{e.parseText(330, 361), e.parseText(362, 393), e.parseText(394, 413), e.parseText(414, 415), e.parseText(416, 424), e.parseText(425, 425), e.parseText(426, 426)}
	contact := Contact{e.parseNumber(315, 324), e.parseText(328, 329), address}
	consumer := Consumer{e.parseText(231, 231), e.parseText(232, 256), e.parseText(257, 276), e.parseText(277, 296), e.parseText(297, 297), e.parseNumber(298, 306), e.parseDate(date, 307, 314), e.parseText(325, 325), e.parseText(326, 327), contact}

	if e.err != nil {
		return nil, e.err
	}

	return &Base{time, correction, identification, cycle, account, consumer}, nil
}
