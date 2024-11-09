package main

import "fmt"
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
*/
func sequence() func() int{
 i:=0
 
 return func() int{
	 i++
     return i
 }
 
}
func sum(numbers ...int) int{
	total:=0
	for _,n :=range numbers{
		total+=n
	}
	return total
}
func main(){
	prices:=[]int{20,31,45,10,15}
	total_price:=sum(prices...)
	fmt.Println("total prices: ", total_price)
	// fmt.Println(processNumbers(10,half))
	nextInt:=sequence()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

	permission("editor")
}