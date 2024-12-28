package service

import "github.com/anton-ag/javacode/internal/repo"

type Wallet interface {
	Update(id string, amount int, operation string) error
	Check(id string) (int, error)
}

type Service struct {
	Wallet Wallet
}

func InitService(repo *repo.Repo) *Service {
	return &Service{
		Wallet: InitWalletService(repo.Wallet),
	}
}
