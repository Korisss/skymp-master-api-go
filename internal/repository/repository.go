package repository

import (
	master_api "github.com/Korisss/skymp-master-api-go"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user master_api.User) (int, error)
	GetUser(email, password string) (master_api.User, error)
	GetUserName(id int) (string, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
