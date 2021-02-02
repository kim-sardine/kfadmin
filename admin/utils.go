package admin

import (
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func getUniqueUUID(uuids []string) string {
	for {
		uuid := uuid.NewV4().String()
		if !contains(uuid, uuids) {
			return uuid
		}
	}
}

func contains(element string, bucket []string) bool {
	for _, v := range bucket {
		if v == element {
			return true
		}
	}
	return false
}

func getUsernameFromEmail(email string) (string, error) {
	idxAt := strings.IndexRune(email, '@')
	if idxAt == -1 {
		return "", fmt.Errorf("wrong email format : %s", email)
	}
	return email[:idxAt], nil
}
