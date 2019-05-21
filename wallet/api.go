package wallet

type (
	RegisterCustomer struct {
		Name        string
		Email       string
		Phonenumber string
	}

	RegisterMerchant struct {
		Name  string
		Email string
	}

	TopUp struct {
		WalletID    string
		Amount      int64
		ReferenceID string
		Description string
	}
)
