package louvois

import (
	"fmt"
	"time"
)

type PayementType struct {
	Id            int
	title         string
	dateDelivery  time.Date
	dateEffective time.Date
}

type Pay struct {
	title     string
	dateStart time.Date
	dateStop  time.Date
}

func GetPay() {
	fmt.Println("coucou")
}
