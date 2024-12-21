package bnk

import "fmt"

type Bank struct {
	name      string
	customers []Customer
}

func (b *Bank) AddCustomers(c ...Customer) {
  for _,cust := range c {
	b.customers = append(b.customers, cust)
  }
}
func (b Bank) DisplayCustomers() {
	fmt.Printf("bank name: %s\n", b.name)
	for _,c:=range b.customers{
           c.DisplayAccounts()
	}

}
func NewBank(name string) *Bank {
	return &Bank{name: name}
}