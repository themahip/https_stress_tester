package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	Handler "stess_tester/handler"
	Types "stess_tester/type"
	"strings"
)

func main() {
	Request := &Types.Request{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the URL you want to test:")
	Request.Url = readInput(reader)

	fmt.Println("Enter the auth token used for user verification:")
	Request.AuthToken = readInput(reader)

	fmt.Println("Enter the HTTP method (POST, GET, PUT, DELETE):")
	Request.Method = readInput(reader)

	fmt.Println("Enter the payload in JSON format (leave blank if none):")
	payloadInput := readInput(reader)
	if payloadInput != "" {
		err := json.Unmarshal([]byte(payloadInput), &Request.Payload)
		if err != nil {
			fmt.Printf("Invalid JSON payload: %v\n", err)
			return
		}
	} else {
		Request.Payload = nil
	}

	fmt.Println("Enter the number of concurrent requests:")
	fmt.Scanln(&Request.ConcurrentRequest)

	fmt.Println("Enter the total number of requests for each concurrent request:")
	fmt.Scanln(&Request.RequestPerUser)

	Handler.Methodhandler(Request)
}

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
