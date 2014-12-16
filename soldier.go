package louvois

import "time"

type Rank struct {
	title     string
	dateStart time.Date
	dateStop  time.Date
}

// amrée air, terre, mere
type Army struct {
	title string
}

type Job struct {
	army            Army
	soldierCategory SoldierCategory
	soldierType     SoldierType
	ranks           Rank
	dateStart       time.Date
	dateStop        time.Date
}

type SoldierCategory struct {
	title string
}
type SoldierType struct {
	title string
}

type SoldierGroup struct {
	title    string
	soldiers []Soldier
}

type Soldier struct {
	firstName     string
	lastName      string
	job           []Job
	adress        []Adress
	phone         []Phone
	missions      []Mission
	vacations     []Vacation
	diseases      []Disease
	responsibleOf []Soldier
	underOrderOf  []Soldier
	hasFamily     bool
	nChildren     int
	dateStartArmy time.Date
	dateBorn      time.Date
	dateDeath     time.Date
}

type Country struct {
	title string
}

type AdressType struct {
	title string
}

type Adress struct {
	adressType   AdressType
	country      Country
	streetName   string
	streetNumber string
	bat          string
	special      string
	dateStart    time.Date
	dateStop     time.Date
}

type PhoneType struct {
	title string
}

type Phone struct {
	number    string
	phoneType PhoneType
}
