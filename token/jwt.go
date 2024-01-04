package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

const MIN_SECRET_KEY_SIZE = 32

type JWTMaker struct {
	secretKey string
}

func NewToken(identity string) (string, *Payload, error) {
	tokenMaker, err := NewJWTMaker(viper.GetString("jwt.symmetricKey"))

	if err != nil {
		return "", nil, err
	}

	return tokenMaker.CreateToken(identity, viper.GetDuration("jwt.accessTokenDuration"))
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < MIN_SECRET_KEY_SIZE {
		return nil, fmt.Errorf("Invalid key size: must be at least %d characters", MIN_SECRET_KEY_SIZE)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(email string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(email, duration)

	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{Issuer: "TODO"}, keyFunc)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, jwt.ErrTokenExpired
		}

		return nil, jwt.ErrTokenMalformed
	}

	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, jwt.ErrTokenMalformed
	}

	return payload, nil
}
