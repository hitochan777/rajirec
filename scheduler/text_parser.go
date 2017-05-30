package scheduler

import (
	"reflect"
	"strconv"
	"log"
)

type Schedule struct {
	Time []int
	Month []int
	Day []int
}

type ParseError struct {
	msg string
}

func (e *ParseError) Error() string {
	return e.msg
}

const (
	Sunday = iota
	Monday = iota
	Tuesday = iota
	Wednesday = iota
	Thursday = iota
	Friday = iota
	Saturday = iota
)

func (sched1 *Schedule) Compare(sched2 *Schedule) bool {
	return reflect.DeepEqual(*sched1, *sched2)
}

func NewSchedule(time []int, month []int, day []int) *Schedule {
	return &Schedule{time, month, day}
}

type Parser struct {
	tokenizer *Tokenizer
	tokens []Token
	schedule *Schedule
}

func NewParser() *Parser {
	tokenizer := NewTokenizer()
	schedule := NewSchedule([]int{}, []int{}, []int{})
	parser := &Parser{tokenizer:tokenizer, schedule:schedule}
	return parser
}

func (p *Parser) getCurrentToken() *Token {
	if len(p.tokens) == 0 {
		return nil
	} else {
		return &p.tokens[0]
	}
}

func (p *Parser) nextToken() {
	if len(p.tokens) > 0 {
		p.tokens = p.tokens[1:]
	}
}

func (p *Parser) Parse(str string) error {
	tokens, err := p.tokenizer.Tokenize(str)
	if err != nil {
		log.Fatal(err)
	}
	return p.parseTokens(tokens)
}

func (p *Parser) parseTokens(tokens []Token) error{
	p.setTokens(tokens)
	return p.parseSchedule()
}

func (p *Parser) parseSchedule() error {
	token := p.getCurrentToken()
	for token != nil {
		switch token.GetTokenCode() {
		case EVERY:
			p.parseEvery()
			break
		case AT:
			p.parseAt()
			break
		case ON:
			p.parseOn()
			break
		default:
			return ParseError{"Failed to Parse"}
			break
		}
		p.nextToken()
		token = p.getCurrentToken()
	}
	return nil
}

func (p *Parser) parseEvery() error {
	token := p.getCurrentToken()
	switch token.GetTokenCode() {
	case WEEKDAY:
		p.schedule.Day = AppendAllIfMissing(p.schedule.Day, []int{1, 2, 3, 4, 5})
		break
	case WEEKEND:
     	p.schedule.Day = AppendAllIfMissing(p.schedule.Day, []int{0, 6})
		break
	default:
		return ParseError{"Failed to Parse"}
		break
	}
	p.nextToken()
	return nil
}

func (p *Parser) parseAt() error {
	token := p.getCurrentToken()
	if token.GetTokenCode() == AT {
		p.nextToken()
		p.parseTime()
	} else {
		return ParseError{"Failed to Parse"}
	}
	p.nextToken()
	return nil
}

func (p *Parser) parseOn() error {
	token := p.getCurrentToken()
	if token.GetTokenCode() == ON {
		p.nextToken()
		p.parseTime()
	} else {
		return ParseError{"Failed to parse"}
	}
	p.nextToken()
	return nil
}

func (p *Parser) parseTime() error {
	token := p.getCurrentToken()
	paramMap := token.GetParamMap()
	hour := paramMap["hour"]
	minute := paramMap["minute"]
	period := paramMap["period"]
	time := 0
	if period == "pm" {
		time += 12 * 3600 // add 12 hours
	}
	if minute != "" {
		if intMinute, err := strconv.Atoi(minute); err != nil {
			time += intMinute * 60
		} else {
			return ParseError{"Failed to parse"}
		}
	}

	if intHour, err := strconv.Atoi(hour); err != nil {
		time += intHour * 3600
	} else {
		return ParseError{"Failed to parse"}
	}
   	p.schedule.Time = AppendIfMissing(p.schedule.Time, time)
	p.nextToken()
	return nil
}

func (p *Parser) parseDay() error {
	token := p.getCurrentToken()
	paramMap := token.GetParamMap()
	weekdayString := paramMap["weekday"]
	var weekday int
	switch weekdayString {
	case "sun", "sunday":
		weekday = 0
		break
	case "mon", "monday":
		weekday = 1
		break
	case "tue", "tuesday":
		weekday = 2
		break
	case "wed", "wednesday"	:
		weekday = 3
		break
	case "thu", "thursday":
		weekday = 4
		break
	case "fri", "friday":
		weekday = 5
		break
	case "sat", "saturday":
		weekday = 6
		break
	default:
		return ParseError{"Failed to parse"}
	}
	p.schedule.Day = AppendIfMissing(p.schedule.Day, weekday)
	return nil
}

func (p *Parser) setTokens(tokens []Token) {
	p.tokens = tokens
}

