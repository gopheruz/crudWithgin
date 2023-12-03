package repo

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	DeletedAt sql.NullTime
}
type UpdateUser struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	DeletedAt sql.NullTime
}
type GetAllUsersParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllUsersResult struct {
	Users []*User
	Count int32
}

type UserStorageI interface {
	Create(u *User) (*User, error)
	Get(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll(params *GetAllUsersParams) (*GetAllUsersResult, error)
	Update(u *UpdateUser) (*User, error)
	Delete(id string) error
}
