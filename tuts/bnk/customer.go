package bnk

import (
	"fmt"
	"time"
)
type Customer struct{
	name string
	accounts []BankAccount
	AuditInfo
}
type AccountNumber string

func (c *Customer) AddAccount(ba BankAccount) {
    c.accounts=append(c.accounts,ba)

}
func (c Customer) DisplayAccounts() {
	fmt.Printf("============ %s account details============ \n",c.name)
	for i,acc:=range c.accounts{
		fmt.Printf("account %d\n",i+1)
        acc.displayBalance()
    }
	
}
// constructor initializing the object with values
func NewCustomer(name string) *Customer {
    return &Customer{name: name,AuditInfo: AuditInfo{time.Now(),time.Now()}}
}