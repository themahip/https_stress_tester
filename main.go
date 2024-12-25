package main

import (
	"fmt"
	Handler "stess_tester/handler"
	Types "stess_tester/type"
)

func main() {
	Request := &Types.Request{}
	// ask url
	fmt.Println("Enter the url you want to test on")
	fmt.Scanln(&Request.Url)

	// ask token
	fmt.Println("enter the auth token used for user verification")
	fmt.Scanln(&Request.AuthToken)

	//ask method
	fmt.Println("enter the Method POST, GET, PUT, Delete")
	fmt.Scanln(&Request.Method)

	fmt.Println("Enter the Payload if any")
	fmt.Scanln(&Request.Payload)

	fmt.Println("Enter the Number of Concurrent Request ")
	fmt.Scanln(&Request.ConcurrentRequest)

	fmt.Println("Enter the Total number of Request for each cuncurrent Request")
	fmt.Scanln(&Request.RequestPerUser)

	Handler.Methodhandler(Request)

}
