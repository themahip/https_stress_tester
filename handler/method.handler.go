package Handler

import (
	"fmt"
	methods "stess_tester/method"
	Types "stess_tester/type"
)

func Methodhandler(Request *Types.Request) {
	switch {
	case Request.Method == "POST":
		methods.POSTMETHOD(Request)
	case Request.Method == "GET":
		methods.GetMethod(Request)
	case Request.Method == "UPDATE":
		fmt.Println(Request.Method)
	case Request.Method == "DELETE":
		fmt.Println(Request.Method)
	default:
		fmt.Println("Invalid Method: ", Request.Method)
	}

}
