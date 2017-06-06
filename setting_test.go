package rajirec

import (
	"testing"
	"reflect"
)

func TestNewConfig(t *testing.T) {
	output := NewConfig("test.yml")
	expected := Config{}
	expected.General.API_URL = "http://www.hoge.com/api.xml"
	expected.DB.DBDir = "dbdir"
	expected.DB.TableName = "tablename"
	expected.DB.DBNAME = "dbname"
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("output: %v\nexpected: %v\n", output, expected)
		return
	}
}

