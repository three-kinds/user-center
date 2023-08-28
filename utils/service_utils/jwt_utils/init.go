package jwt_utils

import (
	"crypto/rsa"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v4"
	"github.com/three-kinds/user-center/initializers"
	"log"
)

func InitJwtUtils() {
	accessTokenPrivateKey = loadPrivateKey(initializers.Config.AccessTokenPrivateKey)
	accessTokenPublicKey = loadPublicKey(initializers.Config.AccessTokenPublicKey)
	accessTokenExpiredDuration = initializers.Config.AccessTokenExpiresIn

	refreshTokenPrivateKey = loadPrivateKey(initializers.Config.RefreshTokenPrivateKey)
	refreshTokenPublicKey = loadPublicKey(initializers.Config.RefreshTokenPublicKey)
	refreshTokenExpiredDuration = initializers.Config.RefreshTokenExpiresIn
}

func loadPrivateKey(base64PrivateKey string) *rsa.PrivateKey {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(base64PrivateKey)
	if err != nil {
		log.Panicln("can not decode private key", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		log.Panicln("can not parse private key", err)
	}
	return key
}

func loadPublicKey(base64PublicKey string) *rsa.PublicKey {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(base64PublicKey)
	if err != nil {
		log.Panicln("can not decode public key", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		log.Panicln("can not parse public key", err)
	}
	return key
}
