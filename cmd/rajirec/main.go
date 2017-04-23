package main

import (
	"github.com/google/subcommands"
	"flag"
	"context"
	"os"

	"github.com/hitochan777/rajirec"
)

func main(){

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(&rajirec.RecordCmd{}, "")
	subcommands.Register(&rajirec.AreaCmd{}, "")
	subcommands.Register(&rajirec.BookCmd{}, "")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
