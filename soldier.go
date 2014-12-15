package louvois

import "time"

type Rank struct {
	title     string
	dateStart time.Date
	dateStop  time.Date
}

type SoldierGroup struct {
	title    string
	soldiers []Soldier
}

type Soldier struct {
	firstName     string
	lastName      string
	ranks         []Rank
	missions      []Mission
	vacations     []Vacation
	diseases      []Disease
	responsibleOf []Soldier
	underOrderOf  []Soldier
	dateBorn      time.Date
	dateDeath     time.Date
}
