package wallet

import "time"

const (
	CUSTOMER = iota
	Merchant
)

type (
	UserAccount struct {
		UserID      string
		Name        string
		Email       string
		Phonenumber string
		UserType    int
	}

	Wallet struct {
		WalletID string
		UserID   string
		Balance  int64
	}

	Transaction struct {
		TransactionID   string
		ReferenceID     string
		CreditWallet    string
		Description     string
		TransactionDate time.Time
		Amount          int64
	}
)

func (w *Wallet) creditBalance(amount int64) {
	w.Balance = + amount
}
