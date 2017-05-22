package rajirec

/******
BNF of Schedule

schedule := every | at | on
every :=

******/


type DayOfWeek int

const (
	SUNDAY DayOfWeek = iota
	MONDAY DayOfWeek = iota
	TUESDAY DayOfWeek = iota
	WEDNESDAY DayOfWeek = iota
	THURSDAY DayOfWeek = iota
	FRIDAY DayOfWeek = iota
	SATURDAY DayOfWeek = iota
)



type Schedule struct {
	Time []int
	Month []int
	Day []DayOfWeekj
}

type Parser struct {
	schedule Schedule
	str []string
	curpos int
}

func (p Parser) Tokenize(str string) {

}

func (p Parser) ParseText(str string) {

}

