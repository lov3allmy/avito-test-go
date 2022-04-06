package service

import (
	"github.com/lov3allmy/avito-test-go/internal/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(repository domain.Repository) domain.Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetUser(userID int) (*domain.User, error) {
	return s.repository.GetUser(userID)
}

func (s *service) CreateUser(user *domain.User) error {
	return s.repository.CreateUser(user)
}

func (s *service) UpdateUser(userID int, user *domain.User) error {
	return s.repository.UpdateUser(userID, user)
}

func (s *service) MakeP2PTransfer(p2pInput domain.P2PInput) error {
	return s.repository.MakeP2PTransfer(p2pInput)
}
