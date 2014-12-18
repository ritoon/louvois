package louvois

import (
	"errors"
	"testing"
)

func TestGetFullName(t *testing.T) {
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
			errors.New("Soldier : nom ou prénom incomplet"),
		},
	}
	for i := 0; i < len(data); i++ {
		s := data[i].in
		res, _ := s.GetFullName()
		if res != data[i].outString {
			t.Error(
				"for ", data[i].title,
				"expected ", data[i].outString,
				"got ", res,
			)
		}

	}
}

func TestSetFullName(t *testing.T) {
	type dataTest struct {
		title  string
		inF    string
		inM    string
		inL    string
		out    string
		outErr error
	}
	data := []dataTest{
		dataTest{
			"A",
			"Marcel",
			"",
			"Bigeard",
			"Marcel Bigeard",
			nil,
		},
		dataTest{
			"B",
			"Marcel",
			"",
			"",
			"",
			errors.New("Soldier : nom ou prénom non renseigné"),
		},
		dataTest{
			"C",
			"Friedrich",
			"von",
			"Wolffhart",
			"Friedrich von Wolffhart",
			nil,
		},
	}
	for i := 0; i < len(data); i++ {
		var s Soldier
		s.SetFullName(data[i].inF, data[i].inM, data[i].inL)
		res, _ := s.GetFullName()
		if res != data[i].out {
			t.Error(
				"for ", data[i].title,
				"expected ", data[i].out,
				"got ", res,
			)
		}

	}
}

func TestSetSocialSecurityNumber(t *testing.T) {
	type dataTest struct {
		title  string
		in     string
		out    string
		outErr error
	}
	data := []dataTest{
		dataTest{
			"A",
			"28812930100143",
			"28812930100143",
			nil,
		},
		dataTest{
			"B",
			"18912930100143",
			"18912930100143",
			errors.New("Soldier : nom ou prénom non renseigné"),
		},
	}
	for i := 0; i < len(data); i++ {
		var s Soldier
		s.SetSocialSecurityNumber(data[i].in)
		res, _ := s.GetSocialSecurityNumber()
		if res != data[i].out {
			t.Error(
				"for ", data[i].title,
				"expected ", data[i].out,
				"got ", res,
			)
		}

	}
}
