package main

import "fmt"

func sayHelloPointer(s *string) {
	fmt.Println(s)
	*s = "hello blaise pointer"

}
func sayHello(s string) {
	s = "hello"
}
func main() {
	// const (
	// 	Agency="Fast Tracks"
	// 	Founded=2001
	// 	Founder="James Carter"
	//     _=iota
	// 	Standard
	// 	FullSize
	// 	// )

	// // 	// fmt.Print(Agency,Standard)
	// //  var name *string
	// // //  fmt.Scanln(&name)
	// // //  fmt.Println("hello",name)
	// // fmt.Println(name)

	// // //  sayHello(name)
	// // //  fmt.Println(name)
	// // //  sayHelloPointer(&name)
	// // //  fmt.Println(name)
	// var   minSpeed,maxSpeed=50,100;
	// speed:=60
	// if speed>minSpeed{
	// 	fmt.Println("good speed")
	// }else if speed<maxSpeed{
	// 	fmt.Println("medium speed")
	// }
	/*
	   hello
	*/
	// isHoliday,isWeekend:=false,false
	// if isHoliday && isWeekend{
	// 	fmt.Println("we are in the holiday of the weekend")
	// } else if !isHoliday && !isWeekend{
	// 	fmt.Println("it is a usual day")
	// }
	var isWeekend bool = true
	var isHoliday bool = false
	isDayoff := isWeekend || isHoliday
	if isDayoff {
		fmt.Println("it is a day off")
	}
	var permission string = ""
	switch permission {
	case "admin":
		fmt.Println("full access ")

	case "editor":
		fmt.Println("public write view")
	case "contributor":
		fmt.Println(" write view")

	default:
		fmt.Println("no permission")

	}
}
