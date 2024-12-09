package main

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func generateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func validateJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Unix() < time.Now().Unix() {
		return nil, fmt.Errorf("token is expired")
	}

	if isTokenRevoked(tokenStr) {
		return nil, fmt.Errorf("token is revoked")
	}

	return claims, nil
}

func isTokenRevoked(token string) bool {
	val, err := RedisClient.Get(Ctx, token).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		return true
	}
	return val == "revoked"
}

func revokeToken(token string, expiry time.Time) error {
	ttl := time.Until(expiry)
	if ttl <= 0 {
		return nil 
	}

	err := RedisClient.Set(Ctx, token, "revoked", ttl).Err()
	if err != nil {
		return err
	}
	return nil
}
