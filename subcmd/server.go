package subcmd

import (
	"flag"
	"context"
	"log"
	"time"

	"github.com/hitochan777/rajirec/db"
	"github.com/robfig/cron"
	"github.com/google/subcommands"
	"os"
	"os/signal"
)

type ServerCmd struct {
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
	c := cron.New()
	scheds := dbm.GetSchedules()
	for _, sched := range scheds {
		jobs := sched.GetCronStrings()
		for _, job := range jobs {
			streamURL, err := areas.GetStreamURL(sched.StationID, sched.Channel)

			if err != nil {
				log.Fatal(err)
			}

			c.AddFunc(job, func() { ServerRecord(streamURL, sched.Prefix, sched.Duration) })
			log.Printf("Registered a schedule %v\n", job)
		}
	}
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	return subcommands.ExitSuccess
}

func ServerRecord(streamURL string, prefix string, duration int) {
    outputPath := prefix + "_" + time.Now().String() + ".m4a"

	log.Printf("Started to record on %s for %d minutes\n" +
		" Recording is saved to %s", streamURL, duration, outputPath)
	Record(streamURL, outputPath, duration)
	log.Printf("Finished recording to %s\n", outputPath)
}
