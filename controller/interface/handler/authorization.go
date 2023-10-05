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
	ContextKey: "userID",
	SigningKey: []byte(os.Getenv("SIGNING_KEY")),
}

func GenerateToken(userID model.ULID) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtConfig.SigningKey)
}

func GetUserID(ctx echo.Context) (model.ULID, error) {
	userID, ok := ctx.Get("userID").(model.ULID)
	if !ok {
		return "", errors.New("invalid token")
	}

	return userID, nil
}
