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
	config := NewConfig(SETTING_FILENAME)
	areas := NewAreas(config.General.API_URL)
	fmt.Println("area\tarea code")
	fmt.Println("=================")
	for _, area := range areas {
		fmt.Printf("%s\t%s\n", area.Areajp, area.Area)
	}
	return subcommands.ExitSuccess
}

type Area struct {
	Areajp string `xml:"areajp"`
	Area string `xml:"area"`
	R1 string `xml:"r1"`
	R2 string `xml:"r2"`
	Fm string `xml:"fm"`
}

type Areas map[string]Area

func NewAreas(configUrl string) Areas {
	areas := struct {
		Areas []Area `xml:"stream_url>data"`
	}{}
	areaMap := Areas{}
	FetchXML(configUrl, &areas)
	for _, area := range areas.Areas {
		areaMap[area.Area] = area
	}
	return areaMap
}
