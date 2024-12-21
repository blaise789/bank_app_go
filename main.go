package main

import (
	"fmt"
	 "github.com/blaise789/bank_app_go/bnk"
)


func main(){
      defer func(){
   r:=recover()
   if r!=nil{
      fmt.Println("Error:",r)
   }
	  }()
	  var bank =*bnk.NewBank("Bank of Kigali")
      var cust1,cust2,cust3 =bnk.NewCustomer("John Doe"),bnk.NewCustomer("Jane Doe"), bnk.NewCustomer("Blaise")
	 acc1:=bnk.NewBankAccount("1234567890",1000)
	 acc2:=bnk.NewBankAccount("9876543210",500)
	 cust1.AddAccount(*acc1)
	 cust1.AddAccount(*acc2)
	 bank.AddCustomers(*cust1,*cust2,*cust3)
	 bank.DisplayCustomers()
	 acc1.Withdraw(400)
	 acc1.Deposit(500)
	

	

    	
 
   
  
}
