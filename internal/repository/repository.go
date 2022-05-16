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

type Verification interface {
	SetDiscord(id string, discord string) error
	SetVerificationCode(id string, code int) error
	GetVerificationCode(id string) (int, error)
}

type Repository struct {
	Authorization
	Verification
}

func NewRepository(db *mongodb.Client) *Repository {
	return &Repository{
		Authorization: mongo.NewAuthMongo(db),
		Verification:  mongo.NewVerificationMongo(db),
	}
}
