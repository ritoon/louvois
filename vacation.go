package louvois

type VacationType struct {
	title string
}

type Vacation struct {
	title          string
	typeOfVacation VacationType
	dateStart      time.Date
	dateStop       time.Stop
}
