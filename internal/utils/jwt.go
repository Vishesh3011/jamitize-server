package utils

import (
	"example/errors"
	"example/internal/types"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func GenerateJWT(userId primitive.ObjectID, email, secret string) (string, errors.AppError) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&jwt.MapClaims{
			"email":   email,
			"user_id": userId,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
			"iat":     time.Now().Unix(),
		},
	)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.ToAppError(err, types.InternalServerError, types.Application)
	}
	return tokenStr, nil
}

func GenerateJWTRefresh(userId, email, secret string) (string, errors.AppError) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&jwt.MapClaims{
			"email":   email,
			"user_id": userId,
			"exp":     time.Now().Add(time.Hour * 360).Unix(),
			"iat":     time.Now().Unix(),
		},
	)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.ToAppError(err, types.InternalServerError, types.Application)
	}
	return tokenStr, nil
}

func VerifyJWT(tokenStr, secret string) errors.AppError {
	tok, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return errors.ToAppError(err, types.BadRequest, types.Application)
	}

	if !tok.Valid {
		err = fmt.Errorf("Invalid token")
		return errors.ToAppError(err, types.BadRequest, types.Controller)
	}
	return nil
}

func ParseJWT(tokenStr, secret string) (*string, *string, errors.AppError) {
	tok, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, nil, errors.ToAppError(err, types.BadRequest, types.Application)
	}

	email := tok.Claims.(jwt.MapClaims)["email"].(string)
	userId := tok.Claims.(jwt.MapClaims)["user_id"].(string)

	if email == "" || userId == "" {
		return nil, nil, errors.ToAppError(err, types.BadRequest, types.Application)
	}

	return &email, &userId, nil
}
