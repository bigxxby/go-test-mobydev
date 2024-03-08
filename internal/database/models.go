package database

import (
	"database/sql"
)

type DatabaseStruct struct {
	DB *sql.DB
}

type User struct {
	Id          int            `json:"id"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Name        sql.NullString `json:"name"`
	Phone       sql.NullString `json:"phone"`
	DateOfBirth sql.NullTime   `json:"date_of_birth"`
	SessionId   sql.NullString `json:"session_id"`
	IsAdmin     sql.NullInt16  `json:"is_admin"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`
}

type UserAP struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
	SessionId   string `json:"session_id"`
	IsAdmin     string `json:"is_admin"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type UpdateUser struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
	IsAdmin     string `json:"is_admin"`
}

type Project struct {
	Id              int            `json:"id"`
	UserId          int            `json:"user_id"`
	Name            sql.NullString `json:"name"`
	Category        sql.NullString `json:"category"`
	ProjectType     sql.NullString `json:"project_type"`
	Year            sql.NullInt32  `json:"year"`
	AgeCategory     sql.NullString `json:"age_category"`
	DurationMinutes sql.NullInt32  `json:"duration_minutes"`
	Keywords        sql.NullString `json:"keywords"`
	Description     sql.NullString `json:"description"`
	Director        sql.NullString `json:"director"`
	Producer        sql.NullString `json:"producer"`
}
