package louvois

import "time"

type MeansType struct {
	title string
}

type MeansCategory struct {
	title string
}

// moyens matériel ou humain
type Means struct {
	title         string
	meansType     MeansType
	meansCategory MeansCategory
}

type Bonus struct {
	title        string
	amount       int
	isPercentage bool
}

type TodoType struct {
	title string
}

// Elements à réaliser
type Todo struct {
	title     string
	todotype  TodoType
	bonus     []Bonus
	state     string
	dateStart time.Time
	dateStop  time.Time
}

type Mission struct {
	title     string
	risk      int
	todos     []Todo
	dateStart time.Time
	dateStop  time.Time
}
