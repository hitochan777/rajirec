package subcmd

import (
	"testing"
	"reflect"
)

func TestNewConfig(t *testing.T) {
	output := NewConfig("test.yml")
	expected := Config{}
	expected.General.API_URL = "http://www3.nhk.or.jp/netradio/app/config_pc_2016.xml"
	expected.DB.Dir = "db"
	expected.DB.Name = "rajirec"
	expected.DB.BookTableName = "book"
	if !reflect.DeepEqual(*output, expected) {
		t.Errorf("output: %v\nexpected: %v\n", *output, expected)
		return
	}
}
