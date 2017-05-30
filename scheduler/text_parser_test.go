package scheduler

import (
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct{
		input string
		expected Schedule
	}{
		{"every weekend", *NewSchedule([]int{}, []int{Sunday, Saturday})},
    	{"every weekday", *NewSchedule([]int{}, []int{Monday, Tuesday, Wednesday, Thursday, Friday})},
    	{"on tue, fri", *NewSchedule([]int{}, []int{Tuesday, Friday})},
    	{"on tue, fri at 20:00", *NewSchedule([]int{72000}, []int{Tuesday, Friday})},
    	{"on tue, fri at 8pm", *NewSchedule([]int{72000}, []int{Tuesday, Friday})},
    	{"at 8pm", *NewSchedule([]int{72000}, []int{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday})},
	}
	p := NewParser()
	for i, test := range tests {
		if err := p.Parse(test.input); err !=nil {
			t.Error(i, err)
		}
		if !p.schedule.Compare(&test.expected) {
			t.Error(i, p.schedule, test.expected)
		}
	}
}
