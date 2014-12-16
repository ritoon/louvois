package louvois

import "time"

type WoundCategory struct {
	title string
}

type WoundType struct {
	title  string
	danger Danger
}

type Wound struct {
	mission       Mission
	woundType     WoundType
	woundCategory WoundCategory
	date          time.Date
}
