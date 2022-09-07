package helpers

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/hamidteimouri/htutils/envier"
	"strconv"
	"strings"
	"time"
)

const (
	bearer       string = "bearer"
	bearerFormat string = "Bearer %s"
)

type JwtToken struct {
	Token string `json:"token"`
}

type JwtClaim struct {
	Name     string `json:"name"`
	Family   string `json:"family"`
	Username string `json:"username"`
	ExpireAt string `json:"expires_at"`
	jwt.StandardClaims
}

func JwtGeneration(name, family, username string) (jwtToken string, err error) {
	exp := envier.EnvOrDefault("JWT_EXPIRE_MINUTES", "60")
	signingKey := envier.Env("JWT_SIGNING_KEY")

	/* string to int */
	_, err = strconv.Atoi(exp)
	if err != nil {
		colog.DoRed("error while convert JWT_EXPIRE_MINUTES to int")
		panic(err)
	}
	expirationTime := time.Now().Add(1 * time.Minute)

	/* preparing data */
	claims := JwtClaim{
		Name:     name,
		Family:   family,
		Username: username,
		ExpireAt: expirationTime.Format(TimeFullFormat),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	/* generate jwt token */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		colog.DoRed("error while generating JWT token")
		return "", err
	}

	return tokenString, nil
}

func JwtTokenValidation(signedToken string) (*JwtClaim, error) {
	signingKey := envier.Env("JWT_SIGNING_KEY")
	token, err := jwt.ParseWithClaims(signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(signingKey), nil
		},
	)

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JwtClaim)

	if !ok {
		err = errors.New("can not parse claims")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return claims, err
	}
	return claims, err
}

func ExtractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], bearer) {
		return "", false
	}

	return authHeaderParts[1], true
}
