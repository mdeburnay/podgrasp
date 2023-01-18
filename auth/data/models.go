package data

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const dbTimeout = time.Second * 3

var db *sql.DB

type Models struct {
	User User
}

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name, omitempty"`
	LastName string `json:"last_name, omitempty"`
	Email string `json:"email, omitempty"`
	Password string `json:"-"`
	Active int `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT id, first_name, last_name, email, active, created_at, updated_at FROM users ORDER by last_name"

	rows, err := db.QueryContext(ctx, query)

	if(err != nil) {
		return nil, err
	}

	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning rows: ", err)
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func New(dbPool * sql.DB) Models {
	db = dbPool
	return Models{
		User :User{},
	}
}

func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT id, first_name, last_name, email, password, active, created_at, updated_at FROM users WHERE email = $1"

	var user User

	row := db.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) GetOne (id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT id, first_name, last_name, email, active, created_at, updated_at FROM users WHERE id = $1"

	var user User
	
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "UPDATE users SET first_name = $1, last_name = $2, email = $3, active = $4, updated_at = $5 WHERE id = $6"

	_, err := db.ExecContext(ctx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Active,
		time.Now(),
		u.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "DELETE FROM users WHERE id = $1"

	_, err := db.ExecContext(ctx, query, u.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) Create() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return 0, err
	}

	var newID int

	query := "INSERT INTO users (first_name, last_name, email, password, active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	err = db.QueryRowContext(ctx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		hashedPassword,
		u.Active,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (u *User) ResetPassword(password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return err
	}

	query := "UPDATE users SET password = $1 WHERE id = $2"

	_, err = db.ExecContext(ctx, query, hashedPassword, u.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) PasswordMatches (plainText string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))

	if err != nil {
		switch {
			case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
				return false, "Invalid password"
			default:
				return false, "Error trying to compare passwords"
		}
	}

	return true, "Password matches"
}
