package rajirec

import (
	"log"
	"flag"
	"github.com/google/subcommands"
	"context"
	"strings"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
	"os"
)

const SETTING_FILENAME = "rajirec.yml"

type SettingCmd struct {
	action string
	key string
}

func (*SettingCmd) Name() string { return "setting" }
func (*SettingCmd) Synopsis() string { return "Change setting" }
func (*SettingCmd) Usage() string {return "rajirec setting"}
func (s *SettingCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&s.action, "action", "", "")
	f.StringVar(&s.key, "key", "", "")
}

func (s *SettingCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	config := NewConfig(SETTING_FILENAME)
	log.Println(s.action)
	switch s.action {
	case "show":
		fmt.Println(config)
		break
	case "set":
		fmt.Errorf("Not implemented") // TODO: implementation
		if s.key == "" {
			log.Println(s.Usage())
			os.Exit(int(subcommands.ExitFailure))
		}
		break
	default:
		log.Println(s.Usage())
		os.Exit(int(subcommands.ExitFailure))
	}
	return subcommands.ExitSuccess
}

func GetKeys(key string) []string {
	return strings.Split(key, ".")
}

type Config struct {
	General struct {
		API_URL string `yaml:"api_url"`
	} `yaml:"general"`

	DB struct {
		Dir string `yaml:"db_dir"`
		Name string `yaml:"db_name"`
		BookTableName string `yaml:"book_table"`
		RecordTableName string `yaml:"record_table"`
	} `yaml:"db"`
}

func NewConfig(fname string) Config {
	config := Config{}
	if raw, err := ioutil.ReadFile(fname); err != nil {
		log.Fatal(err)
	} else {
		if err := yaml.Unmarshal(raw, &config); err != nil{
			log.Fatal(err)
		}
	}
	return config
}
