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
	list bool
}

func (*BookCmd) Name() string { return "book" }
func (*BookCmd) Synopsis() string { return "Book record" }
func (*BookCmd) Usage() string {return "rajirec book"}
func (b *BookCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&b.list, "list", false, "List bookings")
}
func (b *BookCmd) Execute(x context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	config := NewConfig(SETTING_FILENAME)
	if b.list {
		if dbm, err := NewDBManager(config.DB.DBDir, config.DB.DBNAME, config.DB.TableName); err != nil {
			log.Println(err)
			return subcommands.ExitFailure
		} else {
			schedules := dbm.GetSchedules()
			fmt.Printf("%v", schedules)
		}
	} else {
		var date, duration string
		fmt.Print("Date: ")
		fmt.Scanf("%s", &date)
		fmt.Print("Duration: ")
		fmt.Scanf("%s", &duration)
	}
	return subcommands.ExitSuccess
}
