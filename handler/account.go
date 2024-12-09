package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"phase_example/dto/request"
	"phase_example/usecase"
)

type AccountHandler struct {
	usecase usecase.AccountUsecase
}

func (h *AccountHandler) PostAccountSignup(c echo.Context) error {
	req, err := request.NewPostAccountSignup(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := h.usecase.Signup(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func NewAccountHandler(u usecase.AccountUsecase) *AccountHandler {
	return &AccountHandler{
		usecase: u,
	}
}
