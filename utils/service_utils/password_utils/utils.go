package password_utils

import (
	"fmt"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", se.ClientKnownError(fmt.Sprintf("password format error: %s", err))
	}
	return string(hashedPassword), nil
}

func IsSamePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
