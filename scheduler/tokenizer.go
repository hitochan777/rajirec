package scheduler

import (
	"regexp"
	"errors"
	"strconv"
	"log"
	"strings"
)

type TokenCode int

const (
	EVERY TokenCode = iota
	WEEKDAY TokenCode = iota
	WEEKEND TokenCode = iota
	AT TokenCode = iota
	ON TokenCode = iota
	COMMA TokenCode = iota
	TIME TokenCode = iota
	DAY TokenCode = iota
)

type Token struct {
	token TokenCode
	name string
}

func (t Token) GetName() string {
	return t.name
}

func (t Token) GetTokenCode() TokenCode {
	return t.token
}

type TokenInfo struct {
	regex *regexp.Regexp
	token TokenCode
}

func (ti *TokenInfo) FindIndex(str string) []int {
	index := ti.regex.FindIndex([]byte(strings.ToLower(str)))
	return index
}

func NewTokenInfo(regexString string, token TokenCode) *TokenInfo {
	regex, err := regexp.Compile(regexString)
	regex.Longest()
	if err != nil {
		log.Fatal(err)
	}
	tokenInfo := &TokenInfo{regex, token}
	return tokenInfo
}

type Tokenizer struct {
	tokenInfos []TokenInfo
}

func (t *Tokenizer) Tokenize(str string) ([]Token, error) {
	tokens := []Token{}
	for len(str) != 0 {
		match := false
		for _, tokenInfo := range t.tokenInfos {
			index := tokenInfo.FindIndex(str)
			if len(index) != 0 {
				match = true
				tokens = append(tokens, Token{tokenInfo.token, str[index[0]:index[1]]})
				str = str[index[1]:]
			}
		}
		if !match {
			return nil, errors.New("No match!")
		}
	}
	return tokens, nil
}

func (t *Tokenizer) addPattern(pattern string, token TokenCode) {
	for _, tokenInfo := range t.tokenInfos {
		if tokenInfo.token == token {
			log.Fatal("token code " + strconv.Itoa(int(token)) + " already exists in the tokenizer")
		}
	}
	t.tokenInfos = append(t.tokenInfos, *NewTokenInfo(pattern, token))
}

func NewTokenizer() *Tokenizer {
	tokenizer := &Tokenizer{}
	tokenizer.addPattern("^\\s*((([0]?[1-9]|1[0-2]):[0-5]\\d( )?(am|pm))|(([0]?\\d|1\\d|2[0-3]):[0-5]\\d))\\s*", TIME)
	tokenizer.addPattern("^\\s*(every)\\s*", EVERY)
	tokenizer.addPattern("^\\s*(weekday)\\s*", WEEKDAY)
	tokenizer.addPattern("^\\s*(weekend)\\s*", WEEKEND)
	tokenizer.addPattern("^\\s*(at)\\s*", AT)
	tokenizer.addPattern("^\\s*(on)\\s*", ON)
	tokenizer.addPattern("^\\s*(,)\\s*", COMMA)
	tokenizer.addPattern("^\\s*(sun|mon|tue|wed|thu|fri|sat|sunday|monday|tuesday|wednesday|thursday|friday|saturday)\\s*", DAY)
	return tokenizer
}
