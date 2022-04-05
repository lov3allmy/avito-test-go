package user

import "context"

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetUser(ctx context.Context, userID int) (*User, error) {
	return s.repository.GetUser(ctx, userID)
}

func (s *service) CreateUser(ctx context.Context, user *User) error {
	return s.repository.CreateUser(ctx, user)
}

func (s *service) UpdateUser(ctx context.Context, userID int, user *User) error {
	return s.repository.UpdateUser(ctx, userID, user)
}

func (s *service) makeP2PTransfer(ctx context.Context, fromUserID, toUserID, amount int) error {
	return s.repository.makeP2PTransfer(ctx, fromUserID, toUserID, amount)
}
