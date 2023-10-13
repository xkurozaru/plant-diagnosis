package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/xkurozaru/plant-diagnosis/controller/application"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
	"github.com/xkurozaru/plant-diagnosis/controller/interface/messages"
)

type AccountHandler interface {
	SignUp() echo.HandlerFunc
	SignUpAdmin() echo.HandlerFunc
	SignIn() echo.HandlerFunc
	GetUser() echo.HandlerFunc
}

type accountHandler struct {
	accountApplicationService application.AccountApplicationService
}

func NewAccountHandler(
	accountApplicationService application.AccountApplicationService,
) AccountHandler {
	return accountHandler{
		accountApplicationService,
	}
}

func (a accountHandler) SignUp() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req messages.SignUpRequest
		var res messages.SignUpResponse

		err := ctx.Bind(&req)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
		}

		err = a.accountApplicationService.SignUp(req.UserName, req.LoginID, req.Password)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		return ctx.JSON(http.StatusOK, res)
	}
}

func (a accountHandler) SignUpAdmin() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req messages.SignUpRequest
		var res messages.SignUpResponse

		err := ctx.Bind(&req)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
		}

		err = a.accountApplicationService.SignUpAdmin(req.UserName, req.LoginID, req.Password)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		return ctx.JSON(http.StatusOK, res)
	}
}

func (a accountHandler) SignIn() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req messages.SignInRequest
		var res messages.SignInResponse

		err := ctx.Bind(&req)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
		}

		user, err := a.accountApplicationService.SignIn(req.LoginID, req.Password)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		token, err := GenerateToken(user.ID)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}
		res.Token = token

		return ctx.JSON(http.StatusOK, res)
	}
}

func (a accountHandler) GetUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var res messages.GetUserResponse

		userID, err := GetUserID(ctx)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		user, err := a.accountApplicationService.GetUser(model.ULID(userID))
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		res.User = messages.NewUserMessage(user)

		return ctx.JSON(http.StatusOK, res)
	}
}
