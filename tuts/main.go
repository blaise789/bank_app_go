package main

import (
	"fmt"
	"time"
)

// function citizens
/*
func processNumbers(num int,operation func(int) int)int {
	return operation(num)
}
func double(num int) int{
	return num*2
}
func half(num int ) int{
	return num/2
}
// */
// func sequence() func() int{
//  i:=0

//  return func() int{
// 	 i++
//      return i
//  }

// }
// func sum(numbers ...int) int{
// 	total:=0
// 	for _,n :=range numbers{
// 		total+=n
// 	}
// 	return total
// }
func main(){
	// prices:=[]int{20,31,45,10,15}
	// total_price:=sum(prices...)
	// fmt.Println("total prices: ", total_price)
	// // fmt.Println(processNumbers(10,half))
	// nextInt:=sequence()
    // fmt.Println(nextInt())
    // fmt.Println(nextInt())
    // fmt.Println(nextInt())

	// permission("editor")



    defer func(){
		r:=recover();
		if r!=nil{
            fmt.Println("error occurred: ", r)
        }
	}()
	var bankAccount1 *BankAccount=NewBankAccount("rca1234567")
	fmt.Println(bankAccount1)

	
	fmt.Printf("",bankAccount1)
	bankAccount1.displayAccountBalance()
	bankAccount1.deposit(1000)
	bankAccount1.displayAccountBalance()
	errMsg:=bankAccount1.withdraw(500)
	if errMsg!=nil{
        fmt.Println(errMsg)
    }
	bankAccount1.displayAccountBalance()

    var cust1 Customer=Customer{Name: "blaise",AuditInfo: AuditInfo{time.Now(),time.Now()} }
	cust1.addAccount(*bankAccount1)
	cust1.displayAccounts()


}