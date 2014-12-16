package louvois

import "time"

type DiseaseCategory struct {
	title string
}

type DiseaseType struct {
	title string
}

type Disease struct {
	title           string
	diseaseType     DiseaseType
	diseaseCategory DiseaseCategory
	dateStart       time.Date
	dateStop        time.Date
}
