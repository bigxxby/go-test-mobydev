package gotest

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Connection *sql.DB
}

func NewDatabase(dataSourceName string) (*Database, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &Database{Connection: db}, nil
}

func (db *Database) CreateTablesUsers() error {
	tx, err := db.Connection.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL,
            phone TEXT,
            date_of_birth DATE,
			session_id TEXT,
			is_admin INTEGER,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            deleted_at TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) CreateUsersTable() error {
	tx, err := db.Connection.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL,
            phone TEXT,
            date_of_birth DATE,
			session_id TEXT,
			is_admin INTEGER,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            deleted_at TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) CreateProjectsTable() error {
	tx, err := db.Connection.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        CREATE TABLE IF NOT EXISTS projects (
            id INTEGER PRIMARY KEY,
            user_id INTEGER,
            name TEXT,
            category TEXT,
            project_type TEXT,
            year INTEGER,
            age_category TEXT,
            duration_minutes INTEGER,
            keywords TEXT,
            description TEXT,
            director TEXT,
            producer TEXT,
            FOREIGN KEY(user_id) REFERENCES users(id)
        )
    `)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) AddUser(email, password string) error {
	tx, err := db.Connection.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			log.Println("Transaction rolled back due to error:", err.Error())
		}
	}()

	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Println("Error Hashing password : ", err.Error())
		return err
	}

	_, err = tx.Exec(`
        INSERT INTO users (email, password) VALUES (?, ?)
    `, email, string(hashedPassword))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Close() error {
	return db.Connection.Close()
}
