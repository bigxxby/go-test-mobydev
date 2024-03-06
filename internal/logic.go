package gotest

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsValidPassword(password string) bool { // one upper, one string , one num , one special
	if len(password) < 8 {
		return false
	}
	re := regexp.MustCompile(`(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@#$%^&+=!]).*`)
	if !re.MatchString(password) {
		return false
	}

	return true
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
