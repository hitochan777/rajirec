package scheduler

import (
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct{
		input string
		expected Schedule
	}{
		{"every weekend", *NewSchedule([]int{}, []int{}, []int{Sunday, Saturday})},
    	{"every weekday", *NewSchedule([]int{}, []int{}, []int{Monday, Tuesday, Wednesday, Thursday, Friday})},
    	{"on tue, fri", *NewSchedule([]int{}, []int{}, []int{Tuesday, Friday})},
    	{"on tue, fri at 20:00", *NewSchedule([]int{72000}, []int{}, []int{Tuesday, Friday})},
    	{"on tue, fri at 8pm", *NewSchedule([]int{72000}, []int{}, []int{Tuesday, Friday})},
    	{"at 8pm", *NewSchedule([]int{72000}, []int{}, []int{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday})},
	}
	p := NewParser()
	for _, test := range tests {
		schedule := p.Parse(test.input)
		if !schedule.Compare(&test.expected) {
			t.Error(schedule, test.expected)
		}
	}
}
