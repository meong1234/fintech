package wallet

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

type (
	WalletService interface {
		RegisterCustomer(cmd *RegisterCustomer) (string, error)
		RegisterMerchant(cmd *RegisterMerchant) (string, error)
		Topup(cmd *TopUp) (string, error)
	}

	WalletServiceImpl struct {
		userRepo        UserRepo
		walletRepo      WalletRepo
		transactionRepo TransactionRepo
	}
)

func (s *WalletServiceImpl) Topup(cmd *TopUp) (string, error) {
	wallet, err := s.walletRepo.FindByID(cmd.WalletID)
	if err != nil {
		fmt.Printf("error topup on find wallet %s \n", err.Error())
		return "", err
	}

	txnID := uuid.NewV4().String()
	txn := Transaction{
		TransactionID:   txnID,
		ReferenceID:     cmd.ReferenceID,
		CreditWallet:    cmd.WalletID,
		Description:     cmd.Description,
		Amount:          cmd.Amount,
		TransactionDate: time.Now(),
	}

	if err := s.transactionRepo.Save(&txn); err != nil {
		fmt.Printf("error topup %s \n", err.Error())
		return "", err
	}

	wallet.creditBalance(cmd.Amount)
	if err := s.walletRepo.Save(wallet); err != nil {
		fmt.Printf("error topup %s \n", err.Error())
		return "", err
	}

	return txnID, nil
}

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

	if err := s.walletRepo.Save(&wallet); err != nil {
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
		UserID:   userID,
		Name:     cmd.Name,
		Email:    cmd.Email,
		UserType: Merchant,
	}

	return s.createWallet(&account)
}
