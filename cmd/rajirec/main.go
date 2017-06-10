package main

import (
	"github.com/google/subcommands"
	"flag"
	"context"
	"os"

	"github.com/hitochan777/rajirec/subcmd"
)

func main(){
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(&subcmd.RecordCmd{}, "")
	subcommands.Register(&subcmd.AreaCmd{}, "")
	subcommands.Register(&subcmd.BookCmd{}, "")
	subcommands.Register(&subcmd.ServerCmd{}, "")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
