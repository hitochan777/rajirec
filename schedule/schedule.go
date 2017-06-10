package schedule

import (
	"fmt"
	"reflect"

	"github.com/jasonlvhit/gocron"
)

type Schedule struct {
	Time []int
	Day []int
	Duration int
	Channel string
	StationID string
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

func (sched *Schedule) GetCronJobs() []*gocron.Job {
	var jobs []*gocron.Job
	for _, day := range sched.Day {
		for _, time := range sched.Time {
			job := getCronJob(day, time)
			jobs = append(jobs, job)
		}
	}
	return jobs
}

func getCronJob(d int, t int) *gocron.Job {
	var job *gocron.Job
	switch d {
		case Sunday:
			job = gocron.Every(1).Saturday()
			break
		case Monday:
			job = gocron.Every(1).Monday()
			break
		case Tuesday:
			job = gocron.Every(1).Tuesday()
			break
		case Wednesday:
			job = gocron.Every(1).Wednesday()
			break
		case Thursday:
			job = gocron.Every(1).Thursday()
			break
		case Friday:
			job = gocron.Every(1).Friday()
			break
		case Saturday:
			job = gocron.Every(1).Saturday()
			break
		default:
			return nil
	}

	var hours int = t / 3600
	var minutes int = (t % 3600) / 60
	//var seconds int = t % 60

	timeString := fmt.Sprintf("%d:%d", hours, minutes)
	job = job.At(timeString)
	return job
}

func NewSchedule(time []int, day []int) *Schedule {
	return &Schedule{time, day, 0, "", ""}
}
