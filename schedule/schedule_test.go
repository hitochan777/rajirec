package schedule

import (
	"testing"
	"reflect"
)

func TestSchedule_GetCronStrings(t *testing.T) {
	sched := NewSchedule([]int{3600 * 20, 3600 * 22 + 60 * 45}, []int{1, 3})
	expected := []string{"0 0 20 * * 1", "0 45 22 * * 1", "0 0 20 * * 3", "0 45 22 * * 3"}
	cronStrings := sched.GetCronStrings()
	if !reflect.DeepEqual(cronStrings, expected) {
		t.Fatalf("\noutput: %v\nexpected: %v", cronStrings, expected)
	}
}
