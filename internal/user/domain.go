package user

import "context"

type User struct {
	ID      int `json:"id" db:"id"`
	Balance int `json:"balance" db:"balance"`
}

type BalanceOperationInput struct {
	Amount int `json:"amount"`
}

type Repository interface {
	GetUser(ctx context.Context, userID int) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, userID int, user *User) error
	makeP2PTransfer(ctx context.Context, fromUserID, toUserID, amount int) error
}

type Service interface {
	GetUser(ctx context.Context, userID int) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, userID int, user *User) error
	makeP2PTransfer(ctx context.Context, fromUserID, toUserID, amount int) error
}
