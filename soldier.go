package louvois

import (
	"errors"
	"time"
)

type Rank struct {
	title     string
	dateStart time.Time
	dateStop  time.Time
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
	dateStart       time.Time
	dateStop        time.Time
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
	dateStartArmy time.Time
	dateBorn      time.Time
	dateDeath     time.Time
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
	dateStart    time.Time
	dateStop     time.Time
}

type PhoneType struct {
	title string
}

type Phone struct {
	number    string
	phoneType PhoneType
}

func (s *Soldier) GetFullName() (string, error) {
	if s.firstName == "" || s.lastName == "" {
		return "", errors.New("nom ou prénom incomplet")
	}
	return s.firstName + " " + s.lastName, nil
}

func (s *Soldier) SetFullName(f, m, l string) error {
	if f == "" || l == "" {
		return errors.New("nom ou prénom non renseigné")
	}
	s.firstName = f
	if m == "" {
		s.lastName = l
	} else if m == "d'" {
		s.lastName = m + l
	} else {
		s.lastName = m + " " + l
	}

	return nil
}
