package service

import (
	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/Korisss/skymp-master-api-go/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(email, password string) (int, string, error)
	ParseToken(token string) (int, error)
	GetUserName(id int) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
