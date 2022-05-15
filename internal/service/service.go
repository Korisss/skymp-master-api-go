package service

import (
	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/Korisss/skymp-master-api-go/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (string, error)
	GenerateToken(email, password string) (string, string, error)
	ParseToken(token string) (string, error)
	GetUserName(id string) (string, error)
}

type Verification interface {
	SetVerificationCode(id string, code int) error
	GetVerificationCode(id string) (int, error)
}

type Service struct {
	Authorization
	Verification
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Verification:  NewVerificationService(repos.Verification),
	}
}
