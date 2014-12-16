package louvois

import (
	"fmt"
	"time"
)

type PayementType struct {
	Id            int
	title         string
	dateDelivery  time.Time
	dateEffective time.Time
}

type Pay struct {
	title     string
	dateStart time.Time
	dateStop  time.Time
}

func GetPay() {
	fmt.Println("coucou")
}
