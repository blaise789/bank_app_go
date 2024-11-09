package main

import "fmt"

func permission(userRole string) {

	switch userRole {
	case "admin":
		fmt.Println("full system access")
		fallthrough
	case "editor":
		fmt.Println("can publish content")
		fallthrough
	case "contributor":
		fmt.Println("can write content")
		fallthrough
	case "viewer":
		fmt.Println("can view the content")
	default:
		fmt.Println("no permission to our application")
	}

}
