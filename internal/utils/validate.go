package utils

import (
	"regexp"
	"strings"
)

func IsRequired(value string) bool {
	return strings.TrimSpace(value) != ""
}

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

func IsValidPassword(password string) bool {
	return len(password) >= 6
}
