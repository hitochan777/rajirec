package subcmd

import (
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

const SETTING_FILENAME = "rajirec.yml"

type Config struct {
	General struct {
		API_URL string `yaml:"api_url"`
	} `yaml:"general"`

	DB struct {
		Dir string `yaml:"db_dir"`
		Name string `yaml:"db_name"`
		BookTableName string `yaml:"book_table"`
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
