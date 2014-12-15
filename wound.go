package louvois

import "time"

type WoundType struct {
	title  string
	danger Danger
}

type Wound struct {
	mission   Mission
	woundType WoundType
	date      time.Date
}
