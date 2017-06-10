package rajirec

import (
	"testing"
	"reflect"
)

func TestNewConfig(t *testing.T) {
	output := NewConfig("test.yml")
	expected := Config{}
	expected.General.API_URL = "http://www3.nhk.or.jp/netradio/app/config_pc_2016.xml"
	expected.General.OutputDir = "./test/output"
	expected.DB.Dir = "db"
	expected.DB.Name = "rajirec"
	expected.DB.BookTableName = "book"
	expected.DB.RecordTableName = "record"
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("output: %v\nexpected: %v\n", output, expected)
		return
	}
}

