package repo

import (
	"database/sql"
	"fmt"
)

type WalletRepo struct {
	db *sql.DB
}

func (r *WalletRepo) Deposit(id int64, amount int) error {
	query := "UPDATE total SET total + ? WHERE id = ?"

	_, err := r.db.Exec(query, amount, id)
	if err != nil {
		return fmt.Errorf("ошибка пополнения счёта: %w", err)
	}

	return nil
}

func (r *WalletRepo) Withdraw(id int64, amount int) error {
	query := "UPDATE total SET total - ? WHERE id = ?"
	_, err := r.db.Exec(query, amount, id)
	if err != nil {
		return fmt.Errorf("ошибка снятия средств: %w", err)
	}

	return nil
}
