package repo

import (
	"database/sql"
	"fmt"

	"github.com/anton-ag/javacode/internal/domain"
)

type WalletRepo struct {
	db *sql.DB
}

func NewWalletRepo(db *sql.DB) *WalletRepo {
	return &WalletRepo{
		db: db,
	}
}

func (r *WalletRepo) Deposit(id string, amount int) error {
	query := "UPDATE wallet SET total = total + $1 WHERE id = $2"

	res, err := r.db.Exec(query, amount, id)
	if err != nil {
		return fmt.Errorf("ошибка пополнения счёта: %w", err)
	}

	num, _ := res.RowsAffected()
	if num == 0 {
		return domain.ErrWalletNotFound
	}

	return nil
}

func (r *WalletRepo) Withdraw(id string, amount int) error {
	query := "UPDATE wallet SET total = total - $1 WHERE id = $2"
	res, err := r.db.Exec(query, amount, id)
	if err != nil {
		return fmt.Errorf("ошибка снятия средств: %w", err)
	}

	num, _ := res.RowsAffected()
	if num == 0 {
		return domain.ErrWalletNotFound
	}

	return nil
}

func (r *WalletRepo) Check(id string) (int, error) {
	var total int
	query := "SELECT total FROM wallet WHERE id = $1"
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
