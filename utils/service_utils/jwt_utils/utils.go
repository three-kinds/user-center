package jwt_utils

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"time"
)

var accessTokenPrivateKey *rsa.PrivateKey
var accessTokenPublicKey *rsa.PublicKey
var accessTokenExpiredDuration time.Duration
var refreshTokenPrivateKey *rsa.PrivateKey
var refreshTokenPublicKey *rsa.PublicKey
var refreshTokenExpiredDuration time.Duration

func createToken(id int64, expiredDuration time.Duration, key *rsa.PrivateKey) (string, error) {
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = id
	claims["exp"] = now.Add(expiredDuration).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", se.ServerKnownError(fmt.Sprintf("could not generate jwt: %v", err))
	}
	return token, nil
}

func validateToken(token string, key *rsa.PublicKey) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, se.ClientKnownError(fmt.Sprintf("unexpected jwt method: %s", t.Header["alg"]))
		}
		return key, nil
	})

	if err != nil {
		return 0, se.ClientKnownError(fmt.Sprintf("jwt token parse error: %v", err))
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		// todo: se.ValidationError
		return 0, se.ClientKnownError("unexpected jwt claim")
	}

	return int64(claims["sub"].(float64)), nil
}

func CreateAccessToken(id int64) (string, error) {
	return createToken(id, accessTokenExpiredDuration, accessTokenPrivateKey)
}

func ValidateAccessToken(token string) (int64, error) {
	return validateToken(token, accessTokenPublicKey)
}

func CreateRefreshToken(id int64) (string, error) {
	return createToken(id, refreshTokenExpiredDuration, refreshTokenPrivateKey)
}

func ValidateRefreshToken(token string) (int64, error) {
	return validateToken(token, refreshTokenPublicKey)
}
