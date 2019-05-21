package wallet

type (
	UserRepo interface {
		FindByID(userID string) (*UserAccount, error)
		Save(account *UserAccount) (error)
	}

	WalletRepo interface {
		FindByID(walletID string) (*Wallet, error)
		Save(wallet *Wallet) (error)
	}
)
