package postgres

import (
	"ginApi/storage/repo"

	"github.com/jmoiron/sqlx"
)

type USerRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &USerRepo{
		db: db,
	}
}
func (ur *USerRepo) Create(u *repo.User) (*repo.User, error) {
	row := ur.db.QueryRow(`
	INSERT INTO users (id, first_name, last_name, email, created_at, deleted_at)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING *
	`,
		u.ID,
		u.FirstName,
		u.LastName,
		u.Email,
		u.CreatedAt,
		u.DeletedAt,
	)
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.CreatedAt,
		&u.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, err
}

func (ur *USerRepo) Get(id string) (*repo.User, error) {
	var result repo.User
	query := `
		SELECT
	    id,
	    first_name,
	    last_name,
	    email,
	    created_at,
	    deleted_at
	FROM users
	WHERE id = $1 AND deleted_at IS NULL;
	`

	row := ur.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.CreatedAt,
		&result.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *USerRepo) Update(u *repo.UpdateUser) (*repo.User, error) {
	var result repo.User
	query := `
	UPDATE users
	SET
		first_name = $1,
		last_name = $2,
		email = $3
	WHERE id = $4
	RETURNING *
	`
	row := ur.db.QueryRow(query, u.FirstName, u.LastName, u.Email, u.ID)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.CreatedAt,
		&result.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ur *USerRepo) Delete(id string) error {
	query := `
	UPDATE users
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;

	`
	_, err := ur.db.Query(query, id)
	if err != nil {
		return err
	}
	return err
}

func (ur *USerRepo) GetByEmail(email string) (*repo.User, error) {
	var result repo.User
	query := `
			SELECT
			id,
			first_name,
			last_name,
			email,
			created_at,
			deleted_at
		FROM users
		WHERE email = $1 AND deleted_at IS NULL;
	`
	row := ur.db.QueryRow(query, email)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.CreatedAt,
		&result.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, err
}
