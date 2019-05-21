package wallet

import "github.com/stretchr/testify/mock"

type (
	UserRepoMock struct {
		mock.Mock
	}

	WalletRepoMock struct {
		mock.Mock
	}

	TransactionRepoMock struct {
		mock.Mock
	}
)

func (r *TransactionRepoMock) FindByID(txnID string) (*Transaction, error) {
	args := r.Called(txnID)
	return args.Get(0).(*Transaction), args.Error(1)
}

func (r *TransactionRepoMock) Save(txn *Transaction) (error) {
	args := r.Called(txn)
	return args.Error(0)
}

func (r *WalletRepoMock) FindByID(walletID string) (*Wallet, error) {
	args := r.Called(walletID)
	return args.Get(0).(*Wallet), args.Error(1)
}

func (r *WalletRepoMock) Save(wallet *Wallet) (error) {
	args := r.Called(wallet)
	return args.Error(0)
}

func (r *UserRepoMock) FindByID(userID string) (*UserAccount, error) {
	args := r.Called(userID)
	return args.Get(0).(*UserAccount), args.Error(1)
}

func (r *UserRepoMock) Save(account *UserAccount) (error) {
	args := r.Called(account)
	return args.Error(0)
}
