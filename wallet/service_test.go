package wallet

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWalletServiceImpl_RegisterCustomer(t *testing.T) {
	userRepo := &UserRepoMock{}
	walletRepo := &WalletRepoMock{}

	service := WalletServiceImpl{
		userRepo:   userRepo,
		walletRepo: walletRepo,
	}

	cmd := RegisterCustomer{
		Name:        "ApangAIS",
		Email:       "ApangAIS@mail.com",
		Phonenumber: "081108110811",
	}

	userRepo.On("Save", mock.MatchedBy(func(req *UserAccount) bool {
		assert.Equal(t, cmd.Name, req.Name)
		assert.Equal(t, cmd.Email, req.Email)
		assert.Equal(t, cmd.Phonenumber, req.Phonenumber)
		assert.Equal(t, CUSTOMER, req.UserType)
		return true
	})).Return(nil)

	walletRepo.On("Save", mock.Anything).Return(nil)

	walletID, err := service.RegisterCustomer(&cmd)
	assert.NoError(t, err)
	assert.NotNil(t, walletID)
}

func TestWalletServiceImpl_RegisterMerchant(t *testing.T) {
	userRepo := &UserRepoMock{}
	walletRepo := &WalletRepoMock{}

	service := WalletServiceImpl{
		userRepo:   userRepo,
		walletRepo: walletRepo,
	}

	cmd := RegisterMerchant{
		Name:  "ApangAIS",
		Email: "ApangAIS@mail.com",
	}

	userRepo.On("Save", mock.MatchedBy(func(req *UserAccount) bool {
		assert.Equal(t, cmd.Name, req.Name)
		assert.Equal(t, cmd.Email, req.Email)
		assert.Equal(t, MERCHANT, req.UserType)
		return true
	})).Return(nil)

	walletRepo.On("Save", mock.Anything).Return(nil)

	walletID, err := service.RegisterMerchant(&cmd)
	assert.NoError(t, err)
	assert.NotNil(t, walletID)
}

func TestWalletServiceImpl_Topup(t *testing.T) {
	userRepo := &UserRepoMock{}
	walletRepo := &WalletRepoMock{}
	txnRepo := &TransactionRepoMock{}

	service := WalletServiceImpl{
		userRepo:        userRepo,
		walletRepo:      walletRepo,
		transactionRepo: txnRepo,
	}

	walletID := "SOME_WALLET_ID"
	wallet := Wallet{
		WalletID: walletID,
		UserID:   uuid.NewV4().String(),
	}
	walletRepo.On("FindByID", walletID).Return(&wallet, nil)
	walletRepo.On("Save", mock.MatchedBy(func(req *Wallet) bool {
		assert.Equal(t, int64(100000), req.Balance)
		return true
	})).Return(nil)

	cmd := TopUp{
		WalletID:    walletID,
		Description: "FROM BANK XX",
		ReferenceID: "BANK_0001",
		Amount:      int64(100000),
	}

	txnRepo.On("Save", mock.MatchedBy(func(req *Transaction) bool {
		assert.Equal(t, cmd.ReferenceID, req.ReferenceID)
		assert.Equal(t, cmd.Description, req.Description)
		assert.Equal(t, walletID, req.CreditWallet)
		assert.Equal(t, cmd.Amount, req.Amount)
		assert.Equal(t, TXN_TOPUP, req.TransactionType)
		return true
	})).Return(nil)

	txnID, err := service.Topup(&cmd)
	assert.NoError(t, err)
	assert.NotNil(t, txnID)
}

func TestWalletServiceImpl_Pay(t *testing.T) {
	paidAmount := int64(30000)
	customerBalance := int64(50000)

	userRepo := &UserRepoMock{}
	walletRepo := &WalletRepoMock{}
	txnRepo := &TransactionRepoMock{}

	service := WalletServiceImpl{
		userRepo:        userRepo,
		walletRepo:      walletRepo,
		transactionRepo: txnRepo,
	}

	customerWalletID := "SOME_WALLET_ID"
	customerWallet := Wallet{
		WalletID: customerWalletID,
		UserID:   uuid.NewV4().String(),
		Balance:  customerBalance,
	}
	walletRepo.On("FindByID", customerWalletID).Return(&customerWallet, nil)

	merchantUserId := "merchantUserID"
	merchantAccount := UserAccount{
		UserID:   merchantUserId,
		Name:     "STARKOPI",
		Email:    "STARKOPI@mail",
		UserType: MERCHANT,
	}
	userRepo.On("FindMerchantAccountByName", "STARKOPI").Return(&merchantAccount, nil)

	merchantWalletID := "MERCHANT_WALLET_ID"
	merchantWallet := Wallet{
		WalletID: merchantWalletID,
		UserID:   uuid.NewV4().String(),
	}
	walletRepo.On("FindByUserID", merchantUserId).Return(&merchantWallet, nil)

	walletRepo.On("Save", mock.MatchedBy(func(req *Wallet) bool {
		if req.WalletID == merchantWalletID {
			assert.Equal(t, paidAmount, req.Balance)
		}

		if req.WalletID == customerWalletID {
			assert.Equal(t, int64(20000), req.Balance)
		}

		return true
	})).Return(nil)

	cmd := Payment{
		WalletID:    customerWalletID,
		Merchant:    "STARKOPI",
		Description: "PAYMENT FOR STARKOPI XX",
		ReferenceID: "STARKOPI_0001",
		Amount:      paidAmount,
	}

	txnRepo.On("Save", mock.MatchedBy(func(req *Transaction) bool {
		assert.Equal(t, cmd.ReferenceID, req.ReferenceID)
		assert.Equal(t, cmd.Description, req.Description)
		assert.Equal(t, merchantWalletID, req.CreditWallet)
		assert.Equal(t, customerWalletID, req.DebitedWallet)
		assert.Equal(t, cmd.Amount, req.Amount)
		assert.Equal(t, TXN_PAYMENT, req.TransactionType)
		return true
	})).Return(nil)

	txnID, err := service.Pay(&cmd)
	assert.NoError(t, err)
	assert.NotNil(t, txnID)
}
