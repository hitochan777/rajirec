package rajirec

import (
	"fmt"
	"context"

	//"github.com/jasonlvhit/gocron"
	"flag"
	"github.com/google/subcommands"
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
