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
	dateStart       time.Time
	dateStop        time.Time
}

func (d *Disease) GetDiseaseInfos() string {
	var res string
	res = "Maladie : " + d.title + "\n"
	res += "Categorie : " + d.diseaseCategory.title + "\n"
	res += "Type : " + d.diseaseType.title + "\n"
	res += "date début : " + d.dateStart.String() + "\n"
	res += "date fin : " + d.dateStop.String() + "\n"
	return res
}
