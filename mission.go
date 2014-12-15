package louvois

import "time"

type Bonus struct {
	title        string
	amount       int
	isPercentage bool
}

type Mission struct {
	title     string
	risk      int
	bonus     Bonus
	dateStart time.Date
	dateStop  time.Date
}
