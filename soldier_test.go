package louvois

import (
	"errors"
	"testing"
)

func TestFullName(t *testing.T) {
	type dataTest struct {
		title     string
		in        Soldier
		outString string
		outErr    error
	}
	data := []dataTest{
		dataTest{
			"A",
			Soldier{firstName: "Jhon", lastName: "Do"},
			"Jhon Do",
			nil,
		},
		dataTest{
			"B",
			Soldier{firstName: "", lastName: ""},
			"",
			errors.New("nom ou prénom incomplet"),
		},
	}
	for i := 0; i < len(data); i++ {
		s := data[i].in
		res, _ := s.FullName()
		if res != data[i].outString {
			t.Error(
				"for ", data[i].title,
				"expected ", data[i].outString,
				"got ", res,
			)
		}

	}
}
