package dao

import (
	"demo/internal/dao/internal"
)

// walletDao is the data access object for the table wallet.
type walletDao struct {
	*internal.WalletDao
}

var (
	// Wallet is a globally accessible object for table wallet operations.
	Wallet = walletDao{internal.NewWalletDao()}
)
