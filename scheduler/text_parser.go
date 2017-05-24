package scheduler

import (
	"log"
	"reflect"
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
	lookahead Token
}

func NewParser() *Parser {
	tokenizer := NewTokenizer()
	parser := &Parser{tokenizer:tokenizer}
	return parser
}

func (p *Parser) Parse(str string) *Schedule{
	tokens, err := p.tokenizer.Tokenize(str)
	if err != nil {
		log.Fatal(err)
	}
	return p.parseTokens(tokens)
}

func (p *Parser) parseTokens(tokens []Token) *Schedule{
	//TODO: implement
	schedule := &Schedule{}
	return schedule
}

