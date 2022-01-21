package repository

import (
	"fmt"

	master_api "github.com/Korisss/skymp-master-api-go"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user master_api.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash, verified) values ($1, $2, $3, false) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (master_api.User, error) {
	var user master_api.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}

func (r *AuthPostgres) GetUserName(id int) (string, error) {
	var name string

	query := fmt.Sprintf("SELECT name FROM %s WHERE id=$1", usersTable)

	err := r.db.Get(&name, query, id)

	return name, err
}
