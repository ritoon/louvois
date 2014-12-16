package louvois

import (
	"testing"
)

func TestDangerCoef(t *testing.T) {
	type dataTest struct {
		title string
		in    int
		out   float64
	}
	data := []dataTest{
		dataTest{
			"A",
			0,
			0.0,
		},
		dataTest{
			"B",
			1,
			0.1,
		},
		dataTest{
			"C",
			5,
			0.0,
		},
	}
	for i := 0; i < len(data); i++ {
		var d Danger
		d.level = data[i].in
		res := d.dangerCoef()
		if res != data[i].out {
			t.Error(
				"for ", data[i].title,
				"expected ", data[i].out,
				"got ", res,
			)
		}
	}
}
