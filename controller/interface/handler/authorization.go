package handler

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
)

var JwtConfig = echojwt.Config{
	ContextKey: "user",
	SigningKey: []byte(os.Getenv("SIGNING_KEY")),
}

func GenerateToken(userID model.ULID) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID.ToString(),
		"exp":    time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtConfig.SigningKey)
}

func GetUserID(ctx echo.Context) (model.ULID, error) {
	token, ok := ctx.Get(JwtConfig.ContextKey).(*jwt.Token)
	if !ok {
		return model.ULID(""), errors.New("failed to get userID from context")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return model.ULID(""), errors.New("failed to get userID from context")
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return model.ULID(""), errors.New("failed to get userID from context")
	}

	return model.ULID(userID), nil
}
