package bnk

import (
	"errors"
	"fmt"
	"time"
)

type Balance=float64

type  BankAccount struct{
	account_number AccountNumber
	balance float64
	 AuditInfo
}

// deposit
func ( acc *BankAccount) Deposit(amount float64)  error{
	if amount<=0 {
        return errors.New("Invalid deposit amount")
    }
	acc.balance+=amount
	acc.ModifiedAt=time.Now()
	fmt.Println("Deposited: ", amount)
	fmt.Println("Current Balance: ", acc.balance)
	return nil
}

//withdraw
func (acc *BankAccount) Withdraw(amount float64) error{
	if amount<=0 || amount>acc.balance{
        return errors.New("Invalid withdrawal amount or insufficient balance")
    }
    acc.balance-=amount
	acc.ModifiedAt=time.Now()
	fmt.Println("Withdrawn: ", amount)
	fmt.Println("Current Balance: ", acc.balance)
    return nil

}
//displayBalance

func (acc BankAccount) displayBalance(){
	fmt.Printf("    account Number:%s ", acc.account_number)
	fmt.Printf("    balance: %f", acc.balance)
	fmt.Println()
}
//construct bank account
func NewBankAccount(accountNumber AccountNumber,balance Balance ) *BankAccount{
    if!accountNumber.isValid(){
        panic("invalid account number")
    }
    return &BankAccount{account_number: accountNumber,balance: balance,AuditInfo:  AuditInfo{CreatedAt: time.Now(),ModifiedAt: time.Now()}}
}
// validate the accountNumber
func (acc AccountNumber) isValid() bool {
	return len(acc)==10
}