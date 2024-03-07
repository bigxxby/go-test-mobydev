package database

import (
	"gotest/internal/logic"
	"log"
)

func (db *DatabaseStruct) AddUser(email, password string) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			log.Println("Transaction rolled back due to error:", err.Error())
		}
	}()

	hashedPassword, err := logic.HashPassword(password)
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
