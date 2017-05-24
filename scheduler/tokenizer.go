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
				str = strings.TrimSpace(str[index[1]:])
				log.Println(str)
			}
		}
		if !match {
			return nil, errors.New("No match!")
		}
	}
	return tokens, nil
}

func (t *Tokenizer) addPattern(pattern string, token TokenCode) {
	pattern = "^(" + pattern + ")"
	for _, tokenInfo := range t.tokenInfos {
		if tokenInfo.token == token {
			log.Fatal("token code " + strconv.Itoa(int(token)) + " already exists in the tokenizer")
		}
	}
	t.tokenInfos = append(t.tokenInfos, *NewTokenInfo(pattern, token))
}

func NewTokenizer() *Tokenizer {
	tokenizer := &Tokenizer{}
	tokenizer.addPattern("(([0]?[1-9]|1[0-2])(:[0-5]\\d(\\s)?)?(am|pm))|(([0]?\\d|1\\d|2[0-3]):[0-5]\\d)", TIME)
	tokenizer.addPattern("every", EVERY)
	tokenizer.addPattern("weekday", WEEKDAY)
	tokenizer.addPattern("weekend", WEEKEND)
	tokenizer.addPattern("at", AT)
	tokenizer.addPattern("on", ON)
	tokenizer.addPattern(",", COMMA)
	tokenizer.addPattern("sun|mon|tue|wed|thu|fri|sat|sunday|monday|tuesday|wednesday|thursday|friday|saturday", DAY)
	return tokenizer
}
