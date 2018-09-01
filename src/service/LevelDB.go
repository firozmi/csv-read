package service

import (
	"fmt"

	"github.com/hifx/bingo/infra/log"

	"bitbucket.org/firozmi/csv-read/src/conf"
	"github.com/syndtr/goleveldb/leveldb"
)

//DBService represents mysql
type DBService struct {
	dbw  *leveldb.DB
	log  log.Logger
	conf conf.Vars
}

//NewDBService creates a new db service
func NewDBService(conf conf.Vars, log log.Logger) (DBService, error) {
	db, err := leveldb.OpenFile(conf.LevelDB.Path, nil)
	if err != nil {
		return DBService{}, err
	}

	return DBService{dbw: db, log: log, conf: conf}, nil
}

//SaveKeyValue saves key value pair
func (ds DBService) SaveKeyValue(key, val string) error {
	err := ds.dbw.Put([]byte(key), []byte(val), nil)
	return err
}

//GetKeyValue gets value from key
func (ds DBService) GetKeyValue(key string) (val string, err error) {
	value, err := ds.dbw.Get([]byte(key), nil)
	if err != nil {
		return "", err
	}
	val = string(value)
	return val, err
}

//Printall prints all key value pairs in db
func (ds DBService) Printall() {
	iter := ds.dbw.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		ds.log.Info("Printall", fmt.Sprintf("\n key = %s, val = %s", key, value))
	}
	iter.Release()
}
