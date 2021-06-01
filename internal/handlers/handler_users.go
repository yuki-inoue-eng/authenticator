package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuki-inoue-eng/authenticator/internal/models"
)

type UserParam struct {
	LoginID  string `json:"login_id" form:"login_id"`
	Password string `json:"password" form:"password"`
	Qos      string `json:"qos" form:"qos"`
}

func (h *Handler) RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := &UserParam{}
		err := c.Bind(param)
		if err != nil {
			return c.String(http.StatusBadRequest, "failed to bind params.")
		}
		user := models.User{
			LoginID:  param.LoginID,
			Password: param.Password,
		}
		if u, err := h.reps.UserRepository.User(user.LoginID); err != nil {
			return c.String(http.StatusInternalServerError, "there was a problem with db access.")
		} else if u != nil {
			return c.String(http.StatusConflict, fmt.Sprintf("user is already exists. login_id: %s", user.LoginID))
		}
		if err := h.reps.UserRepository.Register(user); err != nil {
			return c.String(http.StatusInternalServerError, "there was a problem with db access.")
		}
		return c.String(http.StatusOK, "")
	}
}
