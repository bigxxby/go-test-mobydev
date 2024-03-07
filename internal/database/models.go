package database

import (
	"database/sql"
)

type DatabaseStruct struct {
	DB *sql.DB
}

type User struct {
	Id          int
	Email       string
	Password    string
	Name        sql.NullString
	Phone       sql.NullString
	DateOfBirth sql.NullTime
	SessionId   sql.NullString
	IsAdmin     sql.NullInt16
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	DeletedAt   sql.NullTime
}

type Project struct {
	Id              int
	UserId          int
	Name            sql.NullString
	Category        sql.NullString
	Project_type    sql.NullString
	Year            sql.NullInt32
	AgeCategory     sql.NullString
	DurationMinutes sql.NullInt32
	Keywords        sql.NullString
	Description     sql.NullString
	Director        sql.NullString
	Producer        sql.NullString
}
