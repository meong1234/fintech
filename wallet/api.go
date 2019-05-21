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
)
