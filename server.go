package rajirec

import (
	"flag"
	"context"
	"github.com/google/subcommands"
	"github.com/jasonlvhit/gocron"
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
	return "server"
}

func (s *ServerCmd) SetFlags(f *flag.FlagSet) {
}

func (s *ServerCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	config := NewConfig(SETTING_FILENAME)
	areas := NewAreas(config.General.API_URL)
	dbm, err := NewDBManager(config.DB.Dir, config.DB.Name, config.DB.BookTableName)
	if err != nil {
		log.Println("No schedule is found. Please book at least once.")
		return subcommands.ExitFailure
	}
	schedules := dbm.GetSchedules()
	for _, schedule := range schedules {
		jobs := schedule.GetCronJobs()
		for _, job := range jobs {
			streamURL, err := areas.GetStreamURL(schedule.StationID, schedule.Channel)
			if err != nil {
				log.Fatal(err)
			}
			job.Do(ServerRecord, streamURL, schedule.Duration)
			log.Printf("Registered a schedule %v\n", job)
		}
	}

	<- gocron.Start()
	return subcommands.ExitSuccess
}

func ServerRecord(streamURL string, duration int) {
	outputPath := GenerateHash()
	log.Printf("Output Path: %s\n", outputPath)
	Record(streamURL, outputPath, duration)
}
