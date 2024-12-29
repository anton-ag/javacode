package domain

import "errors"

var (
	ErrWalletNotFound = errors.New("кошелёк с данным uuid не найден")
)
