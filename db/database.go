package db

import (
    "github.com/nanobox-io/golang-scribble"

	"github.com/hitochan777/rajirec/schedule"
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

func (dbm *DBManager) SaveSchedule(sched schedule.Schedule) error {
	schedules := dbm.GetSchedules()
	schedules = append(schedules, sched)
	if err := dbm.db.Write(dbm.collection, dbm.resource, schedules); err != nil {
		return err
	}
	return nil
}

func (dbm *DBManager) GetSchedules() []schedule.Schedule {
	var scheds []schedule.Schedule
	if err := dbm.db.Read(dbm.collection, dbm.resource, &scheds); err != nil {
		return []schedule.Schedule{}
	}

	return scheds
}

func (dbm *DBManager) GetCollection() string {
	return dbm.collection
}

func (dbm *DBManager) GetResource() string {
	return dbm.resource
}

