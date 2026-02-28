package adapters

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/nathancastelein/go-advanced-features-concurrency/solution/craft/pkg/user"
)

func TestListUsersSQL(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT firstname, lastname FROM users`).WillReturnRows(
		sqlmock.NewRows([]string{"firstname", "lastname"}).
			AddRow("SpongeBob", "SquarePants").
			AddRow("Patrick", "Star"),
	)

	userRepository := &UserSQL{db: db}

	// Act
	users, err := userRepository.List()

	// Assert
	require.NoError(t, err)
	require.Equal(t, []user.User{
		{
			Firstname: "SpongeBob",
			Lastname:  "SquarePants",
		},
		{
			Firstname: "Patrick",
			Lastname:  "Star",
		},
	}, users)
}
