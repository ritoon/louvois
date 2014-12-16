package louvois

import "time"

type DangerType struct {
	title string
}

type DangerCategory struct {
	title string
}

type Danger struct {
	level          int
	dangerType     DangerType
	dangerCategory DangerCategory
	dateStart      time.Time
	dateStop       time.Time
}

func (d *Danger) dangerCoef() float64 {
	switch {
	case d.level == 0:
		return 0.0
	case d.level == 1:
		return 0.1
	}
	return 0.0
}
