package service

import (
	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/Korisss/skymp-master-api-go/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int64, error)
	GenerateToken(email, password string) (int64, string, error)
	ParseToken(token string) (int64, error)
	GetUserName(id int64) (string, error)
}

type Verification interface {
	GetVerificationCode(id int64) (int, error)
	SendCodeToBot(id int64, discord string) error
}

type Service struct {
	Authorization
	// Verification
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		// Verification:  NewVerificationService(repos.Verification),
	}
}
