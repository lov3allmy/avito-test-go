package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lov3allmy/avito-test-go/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_CreateUser(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	r := NewRepository(db)

	type mockBehavior func(user domain.User)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		user         domain.User
		expectedErr  bool
	}{
		{
			name: "OK",
			mockBehavior: func(user domain.User) {
				mock.ExpectExec("INSERT INTO users").WithArgs(user.ID, user.Balance).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			user: domain.User{
				ID:      1,
				Balance: 0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehavior(test.user)
			err := r.CreateUser(&test.user)
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
