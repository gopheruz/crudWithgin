package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"ginApi/models"
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
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, first_name, last_name, email, created_at, deleted_at
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
	RETURNING id, first_name, last_name, email, created_at, deleted_at
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
	var psqlID string
	var psqlEmail string

	rowsId, err := ur.db.Query("SELECT id FROM users WHERE id=$1 AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}
	for rowsId.Next() {
		if err := rowsId.Scan(&psqlID); err != nil {
			return err
		}
	}
	if err := rowsId.Err(); err != nil {
		return err
	}
	if psqlID == "" {
		return errors.New("ID not found in the database maybe alreadey deleted")
	}

	rowsEmail, err := ur.db.Query("SELECT email FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	defer rowsEmail.Close()
	for rowsEmail.Next() {
		if err := rowsEmail.Scan(&psqlEmail); err != nil {
			return err
		}
	}
	if err := rowsEmail.Err(); err != nil {
		return err
	}
	if psqlEmail == "" {
		return errors.New("email not found from database may be already deleted")
	}

	// Perform the update
	query := `
		UPDATE users
		SET deleted_email = $1,
		email = $2,
		deleted_at = CURRENT_TIMESTAMP
		WHERE id = $3;
	`
	_, err = ur.db.Exec(query, psqlEmail, sql.NullString{}, id)
	if err != nil {
		return err
	}

	return nil
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

func (ur *USerRepo) GetAll(params *repo.GetAllUsersParams) (*repo.GetAllUsersResult, error) {
	result := &repo.GetAllUsersResult{
		Users: make([]*models.User, 0),
	}

	offset := (params.Page - 1) * params.Limit
	limit := fmt.Sprintf("LIMIT %d OFFSET %d", params.Limit, offset)
	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter = fmt.Sprintf(`WHERE first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s'`, str, str, str)
	}
	query := `
        SELECT
            id, 
            first_name,
            last_name,
            email,
            created_at
        FROM users ` + filter + ` ORDER BY created_at DESC ` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var resultu models.User
		err := rows.Scan(
			&resultu.ID,
			&resultu.FirstName,
			&resultu.LastName,
			&resultu.Email,
			&resultu.CreatedAt,
		)
		
		if err != nil {
			return nil, err
		}
		result.Users = append(result.Users, &models.User{
			ID: resultu.ID,
			FirstName: resultu.FirstName,
			LastName: resultu.LastName,
			Email: resultu.Email,
			CreatedAt: resultu.CreatedAt,
		})
	}

	queryCount := `SELECT count(1) FROM users ` + filter
	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return result, nil
}
