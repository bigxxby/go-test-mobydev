package logic

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpperCase := regexp.MustCompile(`[A-Z]`).MatchString
	hasLowerCase := regexp.MustCompile(`[a-z]`).MatchString
	hasDigit := regexp.MustCompile(`\d`).MatchString
	hasSpecialChar := regexp.MustCompile(`[@#$%^&+=!]`).MatchString

	return hasUpperCase(password) && hasLowerCase(password) && hasDigit(password) && hasSpecialChar(password)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
