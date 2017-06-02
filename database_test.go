package rajirec

import (
	"testing"
	"path"
	"reflect"
	"os"
	"log"
)

const DB_DIR_NAME = "test_db"
var dbDir string

func TestMain(m *testing.M) {
	var cd string
	var err error
	if cd, err = os.Getwd(); err != nil {
		log.Fatal(err)
	}
	dbDir = path.Join(cd, DB_DIR_NAME)
	if _, err := os.Stat(dbDir); !os.IsNotExist(err) {
		log.Fatal("DB test directory " + dbDir + " already exists. Please delete it first.")
	}

	m.Run() // Run actual tests

	if os.RemoveAll(DB_DIR_NAME); err != nil {
		log.Fatal("Failed to delete " + DB_DIR_NAME)
	}
}

func TestDBMamager_SaveSchedules(t *testing.T) {
	dbm, err := NewDBManager(dbDir, "test", "schedule")
	if err != nil {
		t.Error(err)
		return
	}
	expected := []Schedule{
		{[]int{0, 30},[]int{1, 3}},
		{[]int{15},[]int{0}},
	}
	for _, elem := range expected {
		dbm.SaveSchedule(elem)
	}

	output := dbm.GetSchedules()
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output: %v, Excected: %v\n", output, expected)
	}
}
