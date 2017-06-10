package subcmd

import (
	"flag"
	"context"
	"github.com/google/subcommands"
	"github.com/jasonlvhit/gocron"
	"log"

	"github.com/hitochan777/rajirec/db"
	"time"
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
	config := NewConfig()
	areas := NewAreas(config.General.API_URL)
	dbm, err := db.NewDBManager(config.DB.Dir, config.DB.Name, config.DB.BookTableName)
	if err != nil {
		log.Println("No schedule is found. Please book at least once.")
		return subcommands.ExitFailure
	}
	scheds := dbm.GetSchedules()
	for _, sched := range scheds {
		jobs := sched.GetCronJobs()
		for _, job := range jobs {
			streamURL, err := areas.GetStreamURL(sched.StationID, sched.Channel)
			if err != nil {
				log.Fatal(err)
			}
			outputFile := sched.Prefix + "_" + time.Now().String() + ".m4a"
			job.Do(ServerRecord, streamURL, outputFile, sched.Duration)
			log.Printf("Registered a schedule %v\n", job)
		}
	}

	<- gocron.Start()
	return subcommands.ExitSuccess
}

func ServerRecord(streamURL string, outputPath string, duration int) {
	log.Printf("Started to record on %s for %d minutes\n" +
		" Recording is saved to %s", streamURL, duration, outputPath)
	Record(streamURL, outputPath, duration)
	log.Printf("Finished recording to %s\n", outputPath)
}
