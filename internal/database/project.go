package database

import (
	"log"
)

func (db *DatabaseStruct) AddProject(userId, year, duration int, name, category, projectType, ageCategory, keywords, description, director, producer string) error {
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

	_, err = tx.Exec(`
        INSERT INTO projects (user_id, year, duration_minutes, name, category, project_type, age_category, keywords, description, director, producer)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `, userId, year, duration, name, category, projectType, ageCategory, keywords, description, director, producer)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
