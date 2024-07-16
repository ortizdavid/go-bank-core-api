package entities

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	AccountId 		int64 `gorm:"primaryKey;autoIncrement"`
	CustomerId		int64 `gorm:"column:customer_id"`
	AccountStatus	AccountStatus `gorm:"column:account_status_id"`
	AccountType		AccountType `gorm:"column:account_type_id"`
	AccountNumber	string `gorm:"column:account_number"`
	Iban			string `gorm:"column:iban"`
	HolderName   	string `gorm:"column:holder_name"`
	Balance   		float64 `gorm:"column:balance"`
	Currency   		string `gorm:"column:currency"`
	UniqueId		string `gorm:"column:unique_id"`
	CreatedAt		time.Time `gorm:"column:created_at"`
	UpdatedAt		time.Time `gorm:"column:updated_at"`
	mu 				sync.Mutex
}

func (*Account) TableName() string {
	return "accounts"
}

func (acc *Account) Deposit(amount float64) error {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	err := acc.checkStatus()
	if err != nil {
		return err
	}
	err = acc.validateAmout(amount)
	if err != nil {
		return err
	}
	acc.Balance += amount // operation
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	err := acc.checkStatus()
	if err != nil {
		return err
	}
	err = acc.validateAmout(amount)
	if err != nil {
		return err
	}
	err = acc.checkBalance(amount)
	if err != nil {
		return err
	}
	acc.Balance -= amount // operation
	return nil
}

func (acc *Account) Transfer(amount float64, destination *Account) error {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	destination.mu.Lock()
	defer destination.mu.Unlock()

	err := acc.checkStatus()
	if err != nil {
		return err
	}
	err = destination.checkStatus()
	if err != nil {
		return err
	}
	err = acc.validateAmout(amount)
	if err != nil {
		return err
	}
	acc.Withdraw(amount) // operation
	destination.Deposit(amount) // operation
	return nil
}

func (acc *Account) validateAmout(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("amount must be greather than 0")
	}
	return nil
}

func (acc *Account) checkBalance(amount float64) error {
	if amount > acc.Balance {
		return fmt.Errorf("insufficient founds in account '%s'", acc.AccountNumber)
	}
	return nil
}

func (acc *Account) checkStatus() error {
	if acc.AccountStatus != AccountStatusActive {
		return fmt.Errorf("Account '%s' must be active", acc.AccountNumber)
	}
	return nil
}