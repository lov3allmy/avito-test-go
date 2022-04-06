package domain

//go:generate mockgen -source=domain.go -destination=../mocks/mock.go

type User struct {
	ID      int `json:"id" db:"id"`
	Balance int `json:"balance" db:"balance"`
}

type P2PInput struct {
	FromUserID int `json:"from_user_id" validate:"required"`
	ToUserID   int `json:"to_user_id" validate:"required"`
	Amount     int `json:"amount" validate:"required"`
}

type GetBalanceInput struct {
	ID int `json:"user_id" validate:"required"`
}

type BalanceOperationInput struct {
	UserID int    `json:"user_id" validate:"required"`
	Amount int    `json:"amount" validate:"required"`
	Type   string `json:"type" validate:"required"`
}

type Repository interface {
	GetUser(userID int) (*User, error)
	CreateUser(user *User) error
	UpdateUser(userID int, user *User) error
	MakeP2PTransfer(p2pInput P2PInput) error
}

type Service interface {
	GetUser(userID int) (*User, error)
	CreateUser(user *User) error
	UpdateUser(userID int, user *User) error
	MakeP2PTransfer(p2pInput P2PInput) error
}
