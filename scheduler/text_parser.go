package scheduler

import (
	"log"
	"reflect"
	"time"
)

type Schedule struct {
	Time []int
	Month []int
	Day []int
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

func (p *Parser) Parse(str string) *Schedule{
	tokens, err := p.tokenizer.Tokenize(str)
	if err != nil {
		log.Fatal(err)
	}
	return p.parseTokens(tokens)
}

func (p *Parser) parseTokens(tokens []Token) *Schedule{
	p.setTokens(tokens)
	p.parseSchedule()
}

func (p *Parser) parseSchedule() *Schedule {
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
			log.Fatal("Failed to parse")
			break
		}
		p.nextToken()
		token = p.getCurrentToken()
	}
}

func (p *Parser) parseEvery() {
	token := p.getCurrentToken()
	switch token.GetTokenCode() {
	case WEEKDAY:
		p.schedule.Day = AppendAllIfMissing(p.schedule.Day, []int{1, 2, 3, 4, 5})
		break
	case WEEKEND:
     	p.schedule.Day = AppendAllIfMissing(p.schedule.Day, []int{0, 6})
		break
	default:
		log.Fatal("Failed to parse")
		break
	}
	p.nextToken()
}

func (p *Parser) parseAt() *Schedule {
	token := p.getCurrentToken()
	var sched *Schedule

	if token.GetTokenCode() == AT {
		value := token.GetValue()
		regex := p.tokenizer.tokenInfos[AT].regex
		match := regex.FindStringSubmatch(value)
		paramsMap := make(map[string]string)
		for i, name := range regex.SubexpNames() {
			if i > 0 && i <= len(match) {
				paramsMap[name] = match[i]
			}
		}

	} else {
		log.Fatal("Failed to parse")
	}
	p.nextToken()
	return sched
}

func (p *Parser) parseOn() *Schedule {
	token := p.getCurrentToken()
	var sched *Schedule
	p.nextToken()
	return sched
}

func (p *Parser) parseTime() *Schedule {

}

func (p *Parser) parseDay() *Schedule {

}

func (p *Parser) setTokens(tokens []Token) {
	p.tokens = tokens
}

