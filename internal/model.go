package gotest

import (
	"time"
)

type User struct {
	Id          int
	Name        string
	Email       string
	Password    string
	Phone       string
	SessionId   string
	DateOfBirth time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
