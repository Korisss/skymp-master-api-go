package repository

import (
	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/Korisss/skymp-master-api-go/internal/repository/mongo"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user domain.User) (string, error)
	GetUser(email, password string) (domain.User, error)
	GetUserName(id string) (string, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *mongodb.Client) *Repository {
	return &Repository{
		Authorization: mongo.NewAuthMongo(db),
	}
}
