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
