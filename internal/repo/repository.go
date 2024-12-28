package repo

import "database/sql"

type Wallet interface {
	Deposit(id string, amount int) error
	Withdraw(id string, amount int) error
	Check(id string) (int, error)
}

type Repo struct {
	Wallet
}

func InitRepo(db *sql.DB) *Repo {
	return &Repo{
		Wallet: NewWalletRepo(db),
	}
}
