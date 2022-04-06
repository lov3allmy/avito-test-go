package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lov3allmy/avito-test-go/internal/domain"
)

const (
	QueryGetUser             = "SELECT * FROM users WHERE id = $1"
	QueryCreateUser          = "INSERT INTO users (id, balance) VALUES ($1, $2)"
	QueryUpdateUser          = "UPDATE users SET balance = $1 WHERE id = $2"
	QueryTakeFromUserBalance = "UPDATE users SET balance = (balance - $1) WHERE id = $2"
	QueryPutToUserBalance    = "UPDATE users SET balance = (balance + $1) WHERE id = $2"
)

type repository struct {
	postgres *sqlx.DB
}

func NewRepository(db *sqlx.DB) domain.Repository {
	return &repository{
		postgres: db,
	}
}

func (r *repository) GetUser(userID int) (*domain.User, error) {
	user := &domain.User{}

	err := r.postgres.Get(user, QueryGetUser, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *repository) CreateUser(user *domain.User) error {
	res, err := r.postgres.Exec(QueryCreateUser, user.ID, user.Balance)
	if err != nil {
		return err
	}
	createdRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if createdRows == 0 {
		return errors.New("no one rows created")
	}
	return nil
}

func (r *repository) UpdateUser(userID int, user *domain.User) error {
	res, err := r.postgres.Exec(QueryUpdateUser, user.Balance, userID)
	if err != nil {
		return err
	}
	updatedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if updatedRows == 0 {
		return errors.New("no one rows updated")
	}

	return nil
}

func (r *repository) MakeP2PTransfer(p2pInput domain.P2PInput) error {
	tx, err := r.postgres.Begin()
	if err != nil {
		return err
	}

	res, err := tx.Exec(QueryTakeFromUserBalance, p2pInput.Amount, p2pInput.FromUserID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	updatedRows, err := res.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if updatedRows == 0 {
		_ = tx.Rollback()
		return errors.New("no one rows updated")
	}

	res, err = tx.Exec(QueryPutToUserBalance, p2pInput.Amount, p2pInput.ToUserID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	updatedRows, err = res.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if updatedRows == 0 {
		_ = tx.Rollback()
		return errors.New("no one rows updated")
	}

	return tx.Commit()
}
