package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
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

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		postgres: db,
	}
}

func (r *repository) GetUser(ctx context.Context, userID int) (*User, error) {
	user := &User{}

	err := r.postgres.GetContext(ctx, user, QueryGetUser, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *repository) CreateUser(ctx context.Context, user *User) error {
	res, err := r.postgres.ExecContext(ctx, QueryCreateUser, user.ID, user.Balance)
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

func (r *repository) UpdateUser(ctx context.Context, userID int, user *User) error {
	res, err := r.postgres.ExecContext(ctx, QueryUpdateUser, user.Balance, userID)
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

func (r *repository) makeP2PTransfer(ctx context.Context, fromUserID, toUserID, amount int) error {
	tx, err := r.postgres.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	res, err := tx.Exec(QueryTakeFromUserBalance, amount, fromUserID)
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

	res, err = tx.Exec(QueryPutToUserBalance, amount, toUserID)
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
