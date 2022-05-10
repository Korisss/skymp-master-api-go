package repository

import (
	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/Korisss/skymp-master-api-go/internal/repository/psql"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(email, password string) (domain.User, error)
	GetUserName(id int) (string, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: psql.NewAuthPostgres(db),
	}
}
