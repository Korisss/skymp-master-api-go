package repository

import (
	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/Korisss/skymp-master-api-go/internal/repository/mongo"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user domain.User) (int64, error)
	GetUser(email, password string) (domain.User, error)
	GetUserName(id int64) (string, error)
}

type Verification interface {
	SetDiscord(id int64, discord string) error
	SetVerificationCode(id int64, code int) error
	GetVerificationCode(id int64) (int, error)
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
