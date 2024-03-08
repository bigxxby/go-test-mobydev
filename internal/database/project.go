package database

import (
	"log"
)

func (db *DatabaseStruct) GetAllProjects() ([]Project, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tx, err := db.DB.Begin()
	if err != nil {
		log.Println("Error beginning tx:", err.Error())
		return nil, err
	}

	stmt, err := tx.Prepare("SELECT id, user_id, name, category, project_type, year, age_category, duration_minutes, keywords, description, director, producer FROM projects")
	if err != nil {
		log.Println("Error preparing statement:", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying projects:", err.Error())
		return nil, err
	}
	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var p Project
		err := rows.Scan(
			&p.Id,
			&p.UserId,
			&p.Name,
			&p.Category,
			&p.ProjectType,
			&p.Year,
			&p.AgeCategory,
			&p.DurationMinutes,
			&p.Keywords,
			&p.Description,
			&p.Director,
			&p.Producer,
		)
		if err != nil {
			log.Println("Error scanning project row:", err.Error())
			return nil, err
		}

		projects = append(projects, p)
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error committing tx:", err.Error())
		return nil, err
	}

	return projects, nil
}

func (db *DatabaseStruct) AddProject(userId, year, duration int, name, category, projectType, ageCategory, keywords, description, director, producer string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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
