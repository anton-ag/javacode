package service

import (
	"fmt"

	"github.com/anton-ag/javacode/internal/domain"
	"github.com/anton-ag/javacode/internal/repo"
)

type WalletService struct {
	walletRepo repo.Wallet
}

func InitWalletService(walletRepo repo.Wallet) *WalletService {
	return &WalletService{
		walletRepo: walletRepo,
	}
}

func (s *WalletService) Update(id string, amount int, operation string) error {
	if amount < 0 {
		return fmt.Errorf("указана неверная сумма транзакции")
	}
	switch operation {
	case "DEPOSIT":
		return s.walletRepo.Deposit(id, amount)
	case "WITHDRAW":
		return s.walletRepo.Withdraw(id, amount)
	default:
		return domain.ErrWrongOperation
	}
}

func (s *WalletService) Check(id string) (int, error) {
	return s.walletRepo.Check(id)
}
