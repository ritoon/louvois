package louvois

import (
	"testing"
	"time"
)

func TestGetDiseaseInfos(t *testing.T) {
	type dataTest struct {
		title string
		in    Disease
		out   string
	}
	data := []dataTest{
		dataTest{
			"A",
			Disease{
				title:           "Paludisme",
				diseaseType:     DiseaseType{title: "D"},
				diseaseCategory: DiseaseCategory{title: "P.malariae"},
				dateStart:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				dateStop:        time.Date(2009, time.December, 14, 23, 0, 0, 0, time.UTC),
			},
			"Maladie : Paludisme\nCategorie : P.malariae\nType : D\ndate début : 2009-11-10 23:00:00 +0000 UTC\ndate fin : 2009-12-14 23:00:00 +0000 UTC\n",
		},
	}
	for i := 0; i < len(data); i++ {
		out := data[i].in.GetDiseaseInfos()
		if out != data[i].out {
			t.Error(
				"for ", data[i].title, "\n",
				"expected ", data[i].out, "\n",
				"got ", out,
			)
		}
	}
}
