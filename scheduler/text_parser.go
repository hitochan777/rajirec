package scheduler

type Schedule struct {
	Time []int
	Month []int
	Day []int
}

type Parser struct {
	Tokens []Token
	Lookahead Token
}

func NewParser() *Parser {
	parser := &Parser{}

	return parser
}

func (p Parser) Parse(str string) *Schedule{
	tokenizer := NewTokenizer()
	tokenizer.Tokenize(str)
	//p.Tokens = tokenizer.getTokens()
	return p.ParseTokens()
}

func (p Parser) ParseTokens() *Schedule{
	//TODO: implement
	schedule := &Schedule{}
	return schedule
}

