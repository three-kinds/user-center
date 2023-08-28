package password_utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
	"testing"
)

func TestHashPassword_Success(t *testing.T) {
	password := "123456"
	hashedPassword, err := HashPassword(password)
	assert.Nil(t, err)

	assert.True(t, IsSamePassword(password, hashedPassword))
}

func TestHashPassword_Failed(t *testing.T) {
	password := randstr.String(100)
	_, err := HashPassword(password)
	assert.NotNil(t, err)
	assert.Regexp(t, "password format error", err.Error())
}
