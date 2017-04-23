package rajirec

import (
	"fmt"
	"context"

	"github.com/jasonlvhit/gocron"
	"flag"
	"github.com/google/subcommands"
	"time"
	"strings"
)

type Schedule struct {
	day [7]bool
	at time.Time
	duration time.Duration
}

type BookCmd struct {
	list bool
}

func (*BookCmd) Name() string { return "book" }
func (*BookCmd) Synopsis() string { return "Book record" }
func (*BookCmd) Usage() string {return "rajirec book"}
func (b *BookCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&b.list, "list", false, "List bookings")
}
func (b *BookCmd) Execute(x context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if b.list {
		//TODO: get booking info from database
	} else {
		var date, duration string
		fmt.Print("Date: ")
		fmt.Scanf("%s", &date)
		fmt.Print("Duration: ")
		fmt.Scanf("%s", &duration)


	}
	return subcommands.ExitSuccess
}

func parseStringToFreq(date string, duration_str string) {
	schedule := Schedule{}
	date_split := strings.Split(date, " ")
	for i := range date_split {
		date_split[i] = strings.ToLower(date_split[i])
	}
	switch date_split[0] {
	case "everyday":
		for i := 0; i < 7; i++ {
			schedule.day[i] = true
		}
		break
	}
	for _, val := range date_split {
		switch strings.ToLower(val) {
		case "sunday", "sun":
			schedule.day[0] = true
			break
		case "monday", "mon":
			schedule.day[1] = true
			break
		case "tuesday", "tue":
			schedule.day[2] = true
			break
		case "wednesday", "wed":
			schedule.day[3] = true
			break
		case "thusday", "thu":
			schedule.day[4] = true
			break
		case "friday", "fri":
			schedule.day[5] = true
			break
		case "saturday", "sat":
			schedule.day[6] = true
			break
		default:
			break
		}
	}
	duration, err := time.ParseDuration(duration_str)
	if err != nil {
		panic(err)
	}
	schedule.duration = duration
}
