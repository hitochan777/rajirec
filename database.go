package rajirec

import (
    "github.com/nanobox-io/golang-scribble"
	"log"
)

type DBManager struct {
	collection string
	resource string
	db *scribble.Driver
}

func NewDBManager(collection, resource string) (*DBManager, error) {
	db, err := scribble.New(".", nil)
	dbManager := &DBManager{collection, resource,db}
	return dbManager, err
}

func (dbm *DBManager) SaveSchedule(sched Schedule) bool {
	if err := dbm.db.Write(dbm.collection, dbm.resource, sched); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (dbm *DBManager) GetSchedules() []Schedule {
	scheds := []Schedule{}
	if err := dbm.db.Read(dbm.collection, dbm.resource, &scheds); err != nil {
		log.Fatal("Error", err)
	}
	return scheds
}



