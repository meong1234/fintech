package wallet

import (
	"fmt"
	"github.com/satori/go.uuid"
)


type (
	WalletService interface {
		RegisterCustomer(cmd *RegisterCustomer) (string, error)
		RegisterMerchant(cmd *RegisterMerchant) (string, error)
	}

	WalletServiceImpl struct {
		userRepo UserRepo
		walletRepo WalletRepo
	}
)

func (s *WalletServiceImpl) createWallet(account *UserAccount) (string, error) {
	if err := s.userRepo.Save(account); err != nil {
		fmt.Printf("error creating account %s \n", err.Error())
		return "", err
	}

	walletID := uuid.NewV4().String()
	wallet := Wallet{
		WalletID: walletID,
		UserID:   account.UserID,
	}

	if err :=  s.walletRepo.Save(&wallet); err != nil {
		fmt.Printf("error creating wallet %s \n", err.Error())
		return "", err
	}

	return walletID, nil
}

func (s *WalletServiceImpl) RegisterCustomer(cmd *RegisterCustomer) (string, error) {
	userID := uuid.NewV4().String()

	account := UserAccount{
		UserID:      userID,
		Name:        cmd.Name,
		Email:       cmd.Email,
		Phonenumber: cmd.Phonenumber,
		UserType:    CUSTOMER,
	}

	return s.createWallet(&account)
}

func (s *WalletServiceImpl) RegisterMerchant(cmd *RegisterMerchant) (string, error) {
	userID := uuid.NewV4().String()

	account := UserAccount{
		UserID:      userID,
		Name:        cmd.Name,
		Email:       cmd.Email,
		UserType:    Merchant,
	}

	return s.createWallet(&account)
}

