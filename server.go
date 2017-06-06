package rajirec

import (
	"flag"
	"context"
	"github.com/google/subcommands"
	"github.com/jasonlvhit/gocron"
	"fmt"
	"log"
)

type ServerCmd struct {
	port int
}

func (*ServerCmd) Name() string {
	return "server"
}

func (*ServerCmd) Synopsis() string {
	return "Run server to record"
}

func (*ServerCmd) Usage() string {
	return "TODO: usage"
}

func (s *ServerCmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&s.port, "port", 8080, "port")
}

func (s *ServerCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	dbm, err := NewDBManager("db", "schedule", "schedule")
	if err != nil {
		log.Println("No schedule is found. Please book at least once.")
		return subcommands.ExitFailure
	}
	schedules := dbm.GetSchedules()
	for _, schedule := range schedules {
		jobs := schedule.GetCronJobs()
		for _, job := range jobs {
			job.Do(Record, 1, 2)
		}
	}

	gocron.Every(10).Seconds().Do(func(){
		fmt.Println("Hello")
	})
	<- gocron.Start()
	return subcommands.ExitSuccess
}
