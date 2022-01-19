package service

import "github.com/Korisss/skymp-master-api-go/internal/repository"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
