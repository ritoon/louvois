package louvois

import "time"

type DangerType struct {
	title string
}

type DangerCategory struct {
	title string
}

type Danger struct {
	level int // niveau de danger

	dangerType DangerType // type de danger

	dangerCategory DangerCategory // catégorie

	dateStart time.Time // date début

	dateStop time.Time // date fin
}

/*
  Permet d'ajouter un coef de danger
*/
func (d *Danger) dangerCoef() float64 {
	switch {
	case d.level == 0:
		return 0.0
	case d.level == 1:
		return 0.1
	}
	return 0.0
}
