package subcmd

import (
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"path"
)

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

func GetDefaultConfig() Config {
	conf := Config{}
	conf.General.API_URL = "http://www3.nhk.or.jp/netradio/app/config_pc_2016.xml"
	conf.DB.Dir = path.Join(os.Getenv("HOME"), "rajirec")
	conf.DB.Name = "rajirec"
	conf.DB.BookTableName = "book"
	return conf
}

func NewConfig(fname ...string) *Config {
	if len(fname) >= 2 {
		return nil
	}

	var configFilename string

	if len(fname) == 1 {
		configFilename = fname[0]
	} else if configFilename = os.Getenv("RAJIREC_CONFIG"); configFilename == "" {
		defaultConfig := GetDefaultConfig()
		return &defaultConfig
	}

	config := &Config{}

	if raw, err := ioutil.ReadFile(configFilename); err != nil {
		log.Fatal(err)
	} else {
		if err := yaml.Unmarshal(raw, config); err != nil{
			log.Fatal(err)
		}
	}
	return config
}
