package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Init_DB() (DatabaseStruct, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return DatabaseStruct{}, err
	}

	dbase := DatabaseStruct{
		DB: db,
	}

	return dbase, nil
}

func (db *DatabaseStruct) Close() error {
	return db.DB.Close()
}
