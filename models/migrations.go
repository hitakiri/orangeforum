package models

import (
	"database/sql"
	"log"
)

const ModelVersion = 1

func RunMigrationZero() {
	createConfigTable()
	WriteConfig("version", "1")
}

func Migrate(driverName string, dataSourceName string) error {
	mydb, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	db = mydb

	dbver := DBVersion()
	if dbver > ModelVersion {
		return DBVerAhead
	} else if dbver == ModelVersion {
		return DBMigrationNotNeeded
	}

	for dbver < ModelVersion {
		switch dbver {
		case 0:
			RunMigrationZero()
		}
		newDBVer := DBVersion()
		if newDBVer != dbver + 1 {
			log.Fatal("Migration failed", dbver)
		}
		dbver = newDBVer
	}
	return nil
}

