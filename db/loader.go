package db

import (
	"fmt"

	dbm "github.com/tendermint/tendermint/libs/db"
)

const (
	GoLevelDBBackend = "goleveldb"
	CLevelDBBackend  = "cleveldb"
	MemDBackend      = "memdb"
	BadgerDBBackend  = "badgerdb"
)

type DBWrapper interface {
	dbm.DB
	Compact() error
}

func LoadDB(dbBackend, name, directory string) (DBWrapper, error) {
	switch dbBackend {
	case GoLevelDBBackend:
		return LoadGoLevelDB(name, directory)
	case CLevelDBBackend:
		return LoadCLevelDB(name, directory)
	case MemDBackend:
		return LoadMemDB()
	case BadgerDBBackend:
		return LoadBadgerDB(name,directory)
	default:
		return nil, fmt.Errorf("unknown db backend: %s", dbBackend)
	}
}
