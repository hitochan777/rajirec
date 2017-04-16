package rajirec

import (
	"log"
	"flag"
	"github.com/google/subcommands"
	"os"
	"context"
)

const SETTING_FILENAME = ".rajirec.yaml"

type SettingCmd struct {
	area string
}

func (*SettingCmd) Name() string { return "setting" }
func (*SettingCmd) Synopsis() string { return "Change setting" }
func (*SettingCmd) Usage() string {return "rajirec setting"}
func (s *SettingCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&s.area, "area", "", "duration of recording")
	if s.area == "" {
		log.Fatal("")
		os.Exit(int(subcommands.ExitFailure))
	}
}

func (s *SettingCmd) Execute(x context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if s.area == "" {

	}
	return subcommands.ExitSuccess
}

func GetConfigFilename() string {
	return "http://www3.nhk.or.jp/netradio/app/config_pc_2016.xml"
}
