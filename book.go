package rajirec

import (
	"fmt"
	"context"
	"flag"
	"log"

	"github.com/google/subcommands"
)

type BookCmd struct {
	List bool
	Start string
	Duration int
	StationID string
	Channel string
}


func (*BookCmd) Name() string {
	return "book"
}

func (*BookCmd) Synopsis() string {
	return "Book record"
}

func (*BookCmd) Usage() string {
	return "rajirec book"
}

func (b *BookCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&b.List, "list", false, "list all the bookings")
	f.StringVar(&b.Start, "start", "", "String representing schedule")
	f.IntVar(&b.Duration, "duration", 0, "Duration(minutes)")
	f.StringVar(&b.StationID, "station_id", "", "Station ID")
	f.StringVar(&b.Channel, "channel", "fm", "Channel")
}

func (b *BookCmd) Validate() bool {
	if b.List {
		if b.Start != "" || b.Duration != 0 || b.StationID != "" || b.Channel != "fm" {
			return false
		}
	} else {
		if b.Start == "" || b.Duration == 0 || b.StationID == "" {
			return false
		}
	}
	return true
}

func (b *BookCmd) Execute(x context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if !b.Validate() {
		fmt.Println(b.Usage())
		return subcommands.ExitFailure
	}
	var err error
	var dbm *DBManager
	config := NewConfig(SETTING_FILENAME)
	dbm, err = NewDBManager(config.DB.Dir, config.DB.Name, config.DB.BookTableName)

	if err != nil {
		log.Fatal(err)
	}
	if b.List {
		if dbm, err := NewDBManager(config.DB.Dir, config.DB.Name, config.DB.BookTableName); err != nil {
			log.Println(err)
			return subcommands.ExitFailure
		} else {
			schedules := dbm.GetSchedules()
			fmt.Printf("%v", schedules)
		}
		return subcommands.ExitSuccess
	}

	parser := NewParser()

	err = parser.Parse(b.Start)
	if err != nil {
		log.Fatal(err)
	}
	schedule := parser.GetSchedule()
	schedule.StationID = b.StationID
	schedule.Channel = b.Channel
	schedule.Duration = b.Duration

	log.Printf("Saving a schedule %v", schedule)
	if err := dbm.SaveSchedule(schedule); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully saved")
	log.Println("If you are running the server, you need to restart it to reload the new bookings.")

	return subcommands.ExitSuccess
}
