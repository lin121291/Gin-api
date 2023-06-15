package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execute transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != fn(q) {
		//判斷rollback error
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type Idparams struct {
	Userid int32
}

func (store *Store) DeleteAccount(ctx context.Context, arg Idparams) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		err = store.DeleteList(ctx, arg.Userid)
		if err != nil {
			return err
		}
		err = store.DeleteUsers(ctx, arg.Userid)
		if err != nil {
			return err
		}
		return nil
	})

	return err

}
