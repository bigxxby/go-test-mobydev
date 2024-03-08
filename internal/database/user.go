package database

import (
	"gotest/internal/logic"
	"log"
	"strconv"
	"time"
)

func (db *DatabaseStruct) AddUser(email, password string) error {
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

func (db *DatabaseStruct) UpdateUser(userID string, newName, newPhone string, newDOB string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	now := time.Now()
	_, err = tx.Exec("UPDATE users SET name=?, phone=?, date_of_birth = ?  , updated_at = ? WHERE id=?", newName, newPhone, newDOB, now, userID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *DatabaseStruct) GetAllUsers() ([]UserAP, error) {
	var users []UserAP
	tx, err := db.DB.Begin()
	if err != nil {
		log.Println("Error beginning tx:", err.Error())
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query("SELECT * FROM users")
	if err != nil {
		log.Println("Error querying users table:", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		var userAP UserAP
		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.Name,
			&user.Phone,
			&user.DateOfBirth,
			&user.SessionId,
			&user.IsAdmin,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
		if err != nil {
			log.Println("Error scanning user row:", err.Error())
			continue
		}
		userAP.Id = user.Id
		userAP.Email = user.Email
		userAP.Password = user.Password
		userAP.Name = user.Name.String
		userAP.Phone = user.Phone.String
		userAP.DateOfBirth = user.DateOfBirth.Time.String()
		userAP.SessionId = user.SessionId.String
		userAP.IsAdmin = strconv.Itoa(int(user.IsAdmin.Int16))
		userAP.CreatedAt = user.CreatedAt.Time.String()
		userAP.UpdatedAt = user.UpdatedAt.Time.String()
		if user.DeletedAt.Valid {
			userAP.DeletedAt = user.DeletedAt.Time.String()
		}

		users = append(users, userAP)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over user rows:", err.Error())
		return nil, err
	}

	return users, nil
}

func (db *DatabaseStruct) UpdateUserAdmin(userID string, newName, newPhone, newDOB string, isAdmin bool) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tx, err := db.DB.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer tx.Rollback()

	var isAdminValue int
	if isAdmin {
		isAdminValue = 1
	} else {
		isAdminValue = 0
	}
	now := time.Now()
	_, err = tx.Exec("UPDATE users SET name=?, phone=?, date_of_birth=?, is_admin=?, updated_at=? WHERE id=?",
		newName, newPhone, newDOB, isAdminValue, now, userID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
