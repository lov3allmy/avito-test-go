package domain

//go:generate mockgen -source=domain.go -destination=../mocks/mock.go

type User struct {
	ID      int `json:"id" db:"id"`
	Balance int `json:"balance" db:"balance"`
}

type P2PInput struct {
	FromUserID int `json:"from_user_id" validate:"required,min=0"`
	ToUserID   int `json:"to_user_id" validate:"required,min=0,nefield=FromUserID"`
	Amount     int `json:"amount" validate:"required,min=1"`
}

type GetBalanceInput struct {
	ID int `json:"user_id" validate:"required,min=0"`
}

type BalanceOperationInput struct {
	UserID int    `json:"user_id" validate:"required,min=0"`
	Amount int    `json:"amount" validate:"required,min=1"`
	Type   string `json:"type" validate:"required,oneof=add subtract"`
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
