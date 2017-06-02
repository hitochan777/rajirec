package rajirec

import (
    "github.com/nanobox-io/golang-scribble"
)

type DBManager struct {
	collection string
	resource string
	db *scribble.Driver
}

func NewDBManager(dir, collection, resource string) (*DBManager, error) {
	db, err := scribble.New(dir, nil)
	dbManager := &DBManager{collection, resource,db}
	return dbManager, err
}

func (dbm *DBManager) SaveSchedule(sched Schedule) error {
	schedules := dbm.GetSchedules()
	schedules = append(schedules, sched)
	if err := dbm.db.Write(dbm.collection, dbm.resource, schedules); err != nil {
		return err
	}
	return nil
}

func (dbm *DBManager) GetSchedules() []Schedule {
	scheds := []Schedule{}
	if err := dbm.db.Read(dbm.collection, dbm.resource, &scheds); err != nil {
		return []Schedule{}
	}
	return scheds
}

func (dbm *DBManager) GetCollection() string {
	return dbm.collection
}

func (dbm *DBManager) GetResource() string {
	return dbm.resource
}

