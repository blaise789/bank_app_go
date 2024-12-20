package main

import (
	"errors"
	"fmt"
	"time"
)

type AuditInfo struct {
	CreatedAt  time.Time
	ModifiedAt time.Time
}
type BankAccount struct {
	AccountNumber AccountNumber
	Balance       float64
	AuditInfo
}


func (ba BankAccount) displayAccountBalance() {
	fmt.Println("Current Account Balance: ", ba.AccountNumber, ba.Balance)
}
func (ba *BankAccount) deposit(amount float64) {
	ba.Balance += amount
	ba.ModifiedAt = time.Now()
	fmt.Println("Deposited: ", amount)
	fmt.Println("Current Balance: ", ba.Balance)
}
func (ba *BankAccount) withdraw(amount float64) error {
	if amount == 0 || amount > ba.Balance {
		return errors.New("insufficient balance")
	}
	ba.Balance -= amount
	ba.ModifiedAt = time.Now()
	fmt.Println("Withdrawn: ", amount)
	fmt.Println("Current Balance: ", ba.Balance)
	return nil
}
type Customer struct{
	Name string
	accounts []BankAccount
	AuditInfo
}
func (c *Customer) addAccount(ba BankAccount) {
	c.accounts=append(c.accounts,ba)

}

func (c Customer) displayAccounts(){
	fmt.Printf("customer name %s \n",c.Name)
	for _,acc:=range c.accounts{
		fmt.Println(acc.Balance)
	}
}
func NewCustomer(name string) *Customer {
	return &Customer{Name:name}
}
func NewBankAccount(accountNumber AccountNumber) *BankAccount {
	  if!accountNumber.isValid(){
        panic("invalid account number")
	  }

	return &BankAccount{AccountNumber: accountNumber,Balance: 5000,AuditInfo: AuditInfo{time.Now(),time.Now()}}
}
type AccountNumber string
func (acc AccountNumber) isValid() bool {
	return len(string(acc))==10
}
