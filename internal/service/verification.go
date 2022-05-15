package service

import (
	"github.com/Korisss/skymp-master-api-go/internal/repository"
)

type VerificationService struct {
	repo repository.Verification
}

func NewVerificationService(repo repository.Verification) *VerificationService {
	return &VerificationService{repo: repo}
}

func (s *VerificationService) SetVerificationCode(id string, code int) error {
	return s.repo.SetVerificationCode(id, code)
}

func (s *VerificationService) GetVerificationCode(id string) (int, error) {
	return s.repo.GetVerificationCode(id)
}
