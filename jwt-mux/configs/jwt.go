package configs

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("asdah327tfp2y9f28gy2gp8240gu209")

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
