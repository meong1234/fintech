package wallet

type (
	UserRepo interface {
		FindByID(userID string) (*UserAccount, error)
		FindMerchantAccountByName(name string) (*UserAccount, error)
		Save(account *UserAccount) (error)
	}

	WalletRepo interface {
		FindByID(walletID string) (*Wallet, error)
		FindByUserID(userID string) (*Wallet, error)
		Save(wallet *Wallet) (error)
	}

	TransactionRepo interface {
		FindByID(txnID string) (*Transaction, error)
		Save(txn *Transaction) (error)
	}
)
