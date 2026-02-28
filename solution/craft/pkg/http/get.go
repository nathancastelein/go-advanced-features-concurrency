package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) ListUsers(ctx echo.Context) error {
	users, err := s.user.List()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, users)
}
