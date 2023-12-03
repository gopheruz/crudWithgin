package postgres_test

import (
	"database/sql"
	"testing"
	"time"

	"ginApi/storage/postgres"
	"ginApi/storage/repo"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	userRepo := postgres.NewUser(sqlxDB)

	u := &repo.User{
		ID:        uuid.New(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		DeletedAt: sql.NullTime{},
	}

	// Expectations
	mock.ExpectQuery("^INSERT INTO users").WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "created_at", "deleted_at"}).AddRow(u.ID, u.FirstName, u.LastName, u.Email, u.CreatedAt, u.DeletedAt))

	createdUser, err := userRepo.Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	// Add assertions for the created user data if needed
}

// Similarly, write tests for other functions like Get, Update, Delete, GetByEmail, GetAll
