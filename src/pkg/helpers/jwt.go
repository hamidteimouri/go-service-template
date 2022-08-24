package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/hamidteimouri/htutils/envier"
	"reflect"
	"strconv"
	"time"
)

type JwtToken struct {
	Token string `json:"token"`
}

type JwtClaim struct {
	Name      string `json:"name"`
	Family    string `json:"family"`
	Username  string `json:"username"`
	ExpiresAt string `json:"expires_at"`
	jwt.StandardClaims
}

func JwtGeneration(name, family, username string) (jwtToken string, err error) {
	exp := envier.EnvOrDefault("JWT_EXPIRE_MINUTES", "60")

	// string to int
	min, err := strconv.Atoi(exp)
	if err != nil {
		panic(err)
	}
	fmt.Println("min:", min)
	fmt.Println("type:", reflect.TypeOf(min))
	expirationTime := time.Now().Add(1 * time.Minute)

	/* preparing data */
	claims := JwtClaim{
		Name:      name,
		Family:    family,
		Username:  username,
		ExpiresAt: expirationTime.Format(TimeFullFormat),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	/* generate jwt token */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("supersecretkey"))
	if err != nil {
		colog.DoRed("error while generating JWT token")
		return "", err
	}

	return tokenString, nil
}
