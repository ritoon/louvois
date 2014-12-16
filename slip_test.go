package louvois

import (
	"testing"
)

func TestCreate(t *testing.T) {

	var b Barrack
	b.title = "Casèrne"
	b.adress.country.title = "France"
	b.adress.streetName = "av Maréchal Foch"

	var sol Soldier
	sol.firstName = "Jhonn"
	sol.lastName = "Do"

	var lines []LineSlip

	var s Slip
	s.Create(b, sol, lines)
}
