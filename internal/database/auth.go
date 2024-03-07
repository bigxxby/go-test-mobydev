package database

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (db *DatabaseStruct) CheckUser(email, password string) (*User, bool, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Begin transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, false, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Println("Transaction rolled back due to error:", err.Error())
		}
	}()

	stmt, err := tx.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		return nil, false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(email)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	user := &User{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Phone, &user.DateOfBirth, &user.SessionId, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, false, err
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return nil, false, nil
		}

		// Generate session ID
		sessionID, err := GenerateSessionID()
		if err != nil {
			return nil, false, err
		}

		// Update session ID in the user struct
		user.SessionId.String = sessionID

		// Update session ID in the database
		_, err = tx.Exec("UPDATE users SET session_id = ? WHERE id = ?", sessionID, user.Id)
		if err != nil {
			return nil, false, err
		}

		// Commit transaction
		err = tx.Commit()
		if err != nil {
			log.Println(err.Error())
			return nil, false, err
		}

		return user, true, nil
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		return nil, false, err
	}
	return nil, false, nil
}

func GenerateSessionID() (string, error) {
	sessionID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return sessionID.String(), nil
}

func (db *DatabaseStruct) IsAuthorised(sessionId string) (*User, bool, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Println(err.Error())
		return nil, false, err
	}

	stmt, err := tx.Prepare("SELECT * FROM users WHERE session_id = ?")
	if err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return nil, false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sessionId)
	if err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return nil, false, err
	}
	defer rows.Close()

	user := &User{}

	if rows.Next() {
		if err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Phone, &user.DateOfBirth, &user.SessionId, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
			tx.Rollback()
			log.Println(err.Error())
			return nil, false, err
		}
		log.Println("Logged : ", user)
		if err := tx.Commit(); err != nil {
			log.Println(err.Error())
			return nil, false, err
		}
		return user, true, nil
	}

	if err := tx.Commit(); err != nil {
		log.Println(err.Error())
		return nil, false, err
	}

	return nil, false, nil
}

func (db *DatabaseStruct) LogoutUser(sessionId string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tx, err := db.DB.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Println("Transaction rolled back due to error:2", err.Error())
		}
	}()

	_, err = tx.Exec("UPDATE users SET session_id = '' WHERE session_id = ?", sessionId)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction:", err.Error())
		return err
	}

	return nil
}
