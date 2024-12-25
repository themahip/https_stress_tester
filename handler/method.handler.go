package Handler

import (
	"fmt"
	method "stess_tester/Method"
	Types "stess_tester/type"
)

func Methodhandler(Request *Types.Request) {
	switch {
	case Request.Method == "POST":
		method.GETMETHOD(Request)
	case Request.Method == "GET":
		method.GETMETHOD(Request)
	case Request.Method == "UPDATE":
		fmt.Println(Request.Method)
	case Request.Method == "DELETE":
		fmt.Println(Request.Method)
	default:
		fmt.Println("Invalid Method: ", Request.Method)
	}

}
