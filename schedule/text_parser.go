package schedule

import (
	"strconv"
	"log"

	"github.com/hitochan777/rajirec/util"
)


type ParseError struct {
	msg string
}

func (e *ParseError) Error() string {
	return e.msg
}

type Parser struct {
	tokenizer *Tokenizer
	tokens []Token
	schedule *Schedule
}

func NewParser() *Parser {
	tokenizer := NewTokenizer()
	schedule := NewSchedule([]int{}, []int{})
	parser := &Parser{tokenizer:tokenizer, schedule:schedule}
	return parser
}

func (p *Parser) GetSchedule() Schedule {
	return *p.schedule
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

func (p *Parser) clearSchedule(){
	p.schedule = NewSchedule([]int{}, []int{})
}

func (p *Parser) Parse(str string) error {
	tokens, err := p.tokenizer.Tokenize(str)
	p.clearSchedule()
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
			return &ParseError{"Failed to Parse in parseSchedule"}
			break
		}
		token = p.getCurrentToken()
	}
	if len(p.schedule.Day) == 0 {
		p.schedule.Day = []int{0, 1, 2, 3, 4, 5, 6}
	}
	return nil
}

func (p *Parser) parseEvery() error {
	if p.getCurrentToken().GetTokenCode() != EVERY {
		return &ParseError{"Failed to parse in parseEvery"}
	}
	p.nextToken()
	token := p.getCurrentToken()
	switch token.GetTokenCode() {
	case WEEKDAY:
		p.schedule.Day = util.AppendAllIfMissing(p.schedule.Day, []int{1, 2, 3, 4, 5})
		break
	case WEEKEND:
     	p.schedule.Day = util.AppendAllIfMissing(p.schedule.Day, []int{0, 6})
		break
	default:
		return &ParseError{"Failed to Parse in parseEvery"}
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
		return &ParseError{"Failed to Parse in parseAt"}
	}
	return nil
}

func (p *Parser) parseOn() error {
	token := p.getCurrentToken()
	if token.GetTokenCode() == ON {
		for true {
			p.nextToken()
			p.parseDay()
			if token := p.getCurrentToken(); token == nil || token.GetTokenCode() != COMMA {
				break
			}
		}
	} else {
		return &ParseError{"Failed to parse in parseOn"}
	}
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
		if intMinute, err := strconv.Atoi(minute); err == nil {
			time += intMinute * 60
		} else {
			return &ParseError{"Failed to parse in parseTime"}
		}
	}

	if intHour, err := strconv.Atoi(hour); err == nil {
		time += intHour * 3600
	} else {
		return &ParseError{"Failed to parse in parseTime"}
	}
   	p.schedule.Time = util.AppendIfMissing(p.schedule.Time, time)
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
		return &ParseError{"Failed to parse in parseDay"}
	}
	p.schedule.Day = util.AppendIfMissing(p.schedule.Day, weekday)
	p.nextToken()
	return nil
}

func (p *Parser) setTokens(tokens []Token) {
	p.tokens = tokens
}

