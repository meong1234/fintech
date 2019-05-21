package wallet

import (
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
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
		Name:        "ApangAIS",
		Email:       "ApangAIS@mail.com",
	}

	userRepo.On("Save", mock.MatchedBy(func(req *UserAccount) bool {
		assert.Equal(t, cmd.Name, req.Name)
		assert.Equal(t, cmd.Email, req.Email)
		assert.Equal(t, Merchant, req.UserType)
		return true
	})).Return(nil)

	walletRepo.On("Save", mock.Anything).Return(nil)

	walletID, err := service.RegisterMerchant(&cmd)
	assert.NoError(t, err)
	assert.NotNil(t, walletID)
}
