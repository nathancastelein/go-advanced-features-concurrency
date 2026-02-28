package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/nathancastelein/go-advanced-features-concurrency/solution/craft/pkg/user"
)

var (
	_ user.Lister = &UserListerStub{}
)

type UserListerStub struct{}

func (u *UserListerStub) List() ([]user.User, error) {
	return []user.User{
		{
			Firstname: "SpongeBob",
			Lastname:  "SquarePants",
		},
		{
			Firstname: "Patrick",
			Lastname:  "Star",
		},
	}, nil
}

func TestListUsers(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	server := &Server{user: &UserListerStub{}}

	// Act
	err := server.ListUsers(c)

	// Assert
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	require.JSONEq(t, `[{"Firstname": "SpongeBob", "Lastname": "SquarePants"}, {"Firstname": "Patrick", "Lastname": "Star"}]`, rec.Body.String())
}
