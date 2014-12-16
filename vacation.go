package louvois

import "time"

type VacationType struct {
	title string
}

type Vacation struct {
	title          string
	typeOfVacation VacationType
	dateStart      time.Time
	dateStop       time.Time
}
