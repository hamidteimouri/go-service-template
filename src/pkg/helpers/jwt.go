package helpers

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/hamidteimouri/htutils/htenvier"
	"github.com/sirupsen/logrus"
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
	ID       string `json:"id"`
	ExpireAt string `json:"expires_at"`
	jwt.StandardClaims
}

func JwtGeneration(id string) (jwtToken string, err error) {
	exp := htenvier.EnvOrDefault("JWT_EXPIRE_MINUTES", "60")
	signingKey := htenvier.Env("JWT_SIGNING_KEY")

	/* string to int */
	ex, err := strconv.ParseUint(exp, 10, 64)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("error while convert JWT_EXPIRE_MINUTES to int")
	}
	expirationTime := time.Now().Add(time.Duration(ex) * time.Minute)

	/* preparing data */
	claims := JwtClaim{
		ID:       id,
		ExpireAt: expirationTime.Format(TimeFullFormat),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	/* generate jwt token */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JwtTokenValidation(signedToken string) (*JwtClaim, error) {
	signingKey := htenvier.Env("JWT_SIGNING_KEY")
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
		err = errors.New("can not parse jwt token")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}

	if !token.Valid {
		err = errors.New("token is invalid")
		return nil, err
	}

	return claims, nil
}

// ExtractTokenFromAuthHeader this is a method to get jwt token from authorization header key.
// Inspired from go-kit
func ExtractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], bearer) {
		return "", false
	}

	return authHeaderParts[1], true
}
