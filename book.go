package rajirec

import (
	"fmt"
	"context"

	//"github.com/jasonlvhit/gocron"
	"flag"
	"github.com/google/subcommands"
	"log"
)

type BookCmd struct {
	Action string
	Start string
	Duration int
	StationID string
}

func (*BookCmd) Name() string { return "book" }
func (*BookCmd) Synopsis() string { return "Book record" }
func (*BookCmd) Usage() string {return "rajirec book"}
func (b *BookCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&b.Action, "action", "", "Action to take")
	f.StringVar(&b.Start, "schedule", "", "String representing schedule")
	f.IntVar(&b.Duration, "duration", 0, "duration(minutes)")
	f.StringVar(&b.StationID, "station_id", "", "station ID")
}
func (b *BookCmd) Execute(x context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	var err error
	var dbm *DBManager
	config := NewConfig(SETTING_FILENAME)
	dbm, err = NewDBManager(config.DB.DBDir, config.DB.DBNAME, config.DB.TableName)

	if err != nil {
		log.Fatal(err)
	}
	if b.Action == "list" {
		if dbm, err := NewDBManager(config.DB.DBDir, config.DB.DBNAME, config.DB.TableName); err != nil {
			log.Println(err)
			return subcommands.ExitFailure
		} else {
			schedules := dbm.GetSchedules()
			fmt.Printf("%v", schedules)
		}
		return subcommands.ExitSuccess
	} else if b.Action != "" {
		fmt.Println(b.Usage())
		return subcommands.ExitFailure
	}

	parser := NewParser()

	err = parser.Parse(b.Start)
	if err != nil {
		log.Fatal(err)
	}
	schedule := parser.GetSchedule()
	schedule.StationID = b.StationID
	schedule.Duration = b.Duration

	log.Printf("Saving a schedule %v", schedule)
	if err := dbm.SaveSchedule(schedule); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully saved")

	// Record: TODO: impelement

	return subcommands.ExitSuccess
}
