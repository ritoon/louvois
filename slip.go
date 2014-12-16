package louvois

import "time"

type LineSlipType struct {
	title string
}

type LineSlip struct {
	code         string
	title        string
	number       string
	coef         string
	subTotal     string
	lineSlipType LineSlipType
}

type Slip struct {
	barrack      Barrack
	soldier      Soldier
	lineSlip     []LineSlip
	dateCreation time.Time
	dateEdition  time.Time
}

func (s *Slip) Create(b Barrack, sol Soldier, lines []linesSlip) {

}
