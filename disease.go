package louvois

import "time"

type DiseaseType struct {
	title string
}

type Disease struct {
	title       string
	diseaseType DiseaseType
	dateStart   time.Date
	dateStop    time.Date
}
