package scheduler

import (
	"testing"
)

var tokenInfo_FindIndex_testSet = []struct {
	pattern, str string //input
	tokenCode TokenCode
	expected []int // expected result
 }{
	{"^\\s*every\\s*", "  every weekend", EVERY, []int{0, 8}},
	{"^\\s*every\\s*", "every weekend", EVERY, []int{0, 6}},
    {"every", "abceverye", EVERY, []int{3, 8}},
 }

func TestTokenInfo_FindIndex(t *testing.T) {
	for _, data := range tokenInfo_FindIndex_testSet {
		tokenInfo := NewTokenInfo(data.pattern, data.tokenCode)
		index := tokenInfo.FindIndex(data.str)
		if len(index) == 0 {
			t.Fatal("No Match")
		}
		if index[0] != data.expected[0] || index[1] != data.expected[1] {
			t.Error(index, data.expected)
		}
	}
}

var tokenizerTests = []struct {
	str        string // input
	expected []TokenCode // expected result
}{
	{"every", []TokenCode{EVERY}},
	{"every weekend", []TokenCode{EVERY, WEEKEND}},
	{"every weekend", []TokenCode{EVERY, WEEKEND}},
    {"weekend", []TokenCode{WEEKEND}},
    {"weekday", []TokenCode{WEEKDAY}},
    {"on", []TokenCode{ON}},
    {"at", []TokenCode{AT}},
    {"Wednesday", []TokenCode{DAY}},
    {"wed", []TokenCode{DAY}},
    {"13:00", []TokenCode{TIME}},
    {"13::010pm", nil},
}

func TestTokenizer(t *testing.T){
	var tokens []Token
	tok := NewTokenizer()
	for _, data := range tokenizerTests {
		tokens, _ = tok.Tokenize(data.str)
		if data.expected == nil {
			if tokens != nil {
				t.Error("expected nil but output is not nil")
			}
		}
		if len(tokens) != len(data.expected) {
			//t.Error("expected: " + string(data.expected) + "\noutput: " + string(tokens))
			t.Error()
			continue
		}
		for i, token := range data.expected {
			output := tokens[i].GetTokenCode()
			if output != token {
				//t.Error("expected: " + string(data.expected) + "\noutput: " + string(tokens))
				t.Error()
			}
		}
	}
}
