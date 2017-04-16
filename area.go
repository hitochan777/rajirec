package rajirec

import (
	"flag"
	"context"
	"fmt"

	"github.com/google/subcommands"
)

type AreaCmd struct {}

func (*AreaCmd) Name() string { return "area" }
func (*AreaCmd) Synopsis() string { return "Show area information" }
func (*AreaCmd) Usage() string {return "rajirec area"}
func (r *AreaCmd) SetFlags(f *flag.FlagSet) {}
func (r *AreaCmd) Execute(x context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	areas := NewAreas(GetConfigFilename())
	fmt.Println("area\tarea code")
	fmt.Println("=================")
	for _, area := range areas {
		fmt.Printf("%s\t%s\n", area.Areajp, area.Area)
	}
	return subcommands.ExitSuccess
}
