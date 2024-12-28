package repo

import (
	"database/sql"
	"fmt"
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
	query := "UPDATE total SET total + ? WHERE id = ?"

	_, err := r.db.Exec(query, amount, id)
	if err != nil {
		return fmt.Errorf("ошибка пополнения счёта: %w", err)
	}

	return nil
}

func (r *WalletRepo) Withdraw(id string, amount int) error {
	query := "UPDATE total SET total - ? WHERE id = ?"
	_, err := r.db.Exec(query, amount, id)
	if err != nil {
		return fmt.Errorf("ошибка снятия средств: %w", err)
	}

	return nil
}

func (r *WalletRepo) Check(id string) (int, error) {
	var total int
	query := "SELECT total WHERE id = ?"
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
