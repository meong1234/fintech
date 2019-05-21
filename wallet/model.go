package wallet

const (
	CUSTOMER = iota
	Merchant
)

type (
	UserAccount struct {
		UserID string
		Name string
		Email string
		Phonenumber string
		UserType int
	}

	Wallet struct {
		WalletID string
		UserID string
	}
)

