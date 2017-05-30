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
	value string
	tokenCode TokenCode
	paramMap map[string]string
}

func (t *Token) GetParamMap() map[string]string {
	return t.paramMap
}

func (t *Token) GetTokenCode() TokenCode {
	return t.tokenCode
}

func (t *Token) GetValue() string {
	return t.value
}

type TokenInfo struct {
	regex *regexp.Regexp
	tokenCode TokenCode
}

func (ti *TokenInfo) GetTokenCode() TokenCode {
	return ti.tokenCode
}

func (ti *TokenInfo) FindIndex(str string) []int {
	index := ti.regex.FindIndex([]byte(strings.ToLower(str)))
	return index
}

func (ti *TokenInfo) GetParams(str string) (paramsMap map[string]string) {
	match := ti.regex.FindStringSubmatch(str)
	paramsMap = make(map[string]string)
	for i, name := range ti.regex.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
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
	tokenInfos map[TokenCode]TokenInfo
}

func (t *Tokenizer) Tokenize(str string) ([]Token, error) {
	tokens := []Token{}
	for len(str) != 0 {
		match := false
		for tokenCode, tokenInfo := range t.tokenInfos {
			index := tokenInfo.FindIndex(str)
			paramMap := tokenInfo.GetParams(str)
			if len(index) != 0 {
				match = true
				tokens = append(tokens, Token{paramMap: paramMap, tokenCode: tokenCode, value:str[index[0]:index[1]]})
				str = strings.TrimSpace(str[index[1]:])
			}
		}
		if !match {
			return nil, errors.New("No match!")
		}
	}
	return tokens, nil
}

func (t *Tokenizer) addPattern(pattern string, tokenCode TokenCode) {
	pattern = "^(" + pattern + ")"
	if _, ok := t.tokenInfos[tokenCode]; !ok {
		log.Fatal("token code " + strconv.Itoa(int(tokenCode)) + " already exists in the tokenizer")
	}
	t.tokenInfos[tokenCode] = *NewTokenInfo(pattern, tokenCode)
}

func NewTokenizer() *Tokenizer {
	tokenizer := &Tokenizer{}
	tokenizer.addPattern("((?P<hour>[0]?[1-9]|1[0-2])(:(?<minute>[0-5]\\d)(\\s)?)?(?<period>am|pm))|((?<hour>[0]?\\d|1\\d|2[0-3]):(?<minute>[0-5]\\d))", TIME)
	tokenizer.addPattern("every", EVERY)
	tokenizer.addPattern("weekday", WEEKDAY)
	tokenizer.addPattern("weekend", WEEKEND)
	tokenizer.addPattern("at", AT)
	tokenizer.addPattern("on", ON)
	tokenizer.addPattern(",", COMMA)
	tokenizer.addPattern("(?P<weekday>sun|mon|tue|wed|thu|fri|sat|sunday|monday|tuesday|wednesday|thursday|friday|saturday)", DAY)
	return tokenizer
}
