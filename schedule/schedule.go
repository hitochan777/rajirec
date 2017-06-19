package schedule

import (
	"fmt"
	"reflect"
)

type Schedule struct {
	Time []int
	Day []int
	Duration int
	Channel string
	StationID string
	Prefix string
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

func (sched *Schedule) GetCronStrings() []string {
	var jobs []string
	for _, day := range sched.Day {
		for _, time := range sched.Time {
			job := getCronString(day, time)
			jobs = append(jobs, job)
		}
	}
	return jobs
}

func getCronString(d int, t int) string {
	var hour int = t / 3600
	var minute int = (t % 3600) / 60
	//var seconds int = t % 60

	timeString := fmt.Sprintf("0 %d %d * * %d", minute, hour, d)
	return timeString
}

func NewSchedule(time []int, day []int) *Schedule {
	return &Schedule{Time: time, Day:day}
}
