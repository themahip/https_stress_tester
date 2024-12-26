package result

import (
	"fmt"
	Types "stess_tester/type"
	"time"
)

func Processresult(perresults <-chan Types.PerResult, request *Types.Request) {
	var SuccessRequest, FailedRequest int
	var totalDuration time.Duration
	// var errormessage []string

	for perresult := range perresults {
		if perresult.Err != nil {
			FailedRequest++
			// errormessage = append(errormessage, perresult.Err.Error())
			totalDuration = +perresult.Duration
		} else {
			SuccessRequest++
			totalDuration = +perresult.Duration
		}
	}

	finalResult := Types.FinalResult{
		TargetUrl:             request.Url,
		TotalRequestSent:      request.RequestPerUser * request.ConcurrentRequest,
		ConcurrentRequestSent: request.ConcurrentRequest,
		TotalDuration:         totalDuration,
		FailedRequest:         FailedRequest,
		SuccessRequest:        SuccessRequest,
	}
	fmt.Println("Target Url: ", finalResult.TargetUrl)
	fmt.Println("Total Request Sent: ", finalResult.TotalRequestSent)
	fmt.Println("Total Concurrent request sent: ", finalResult.ConcurrentRequestSent)
	fmt.Println("Total Duration: ", finalResult.TotalDuration)
	fmt.Println("Failed Request: ", finalResult.FailedRequest)
	fmt.Println("Success Request: ", finalResult.SuccessRequest)
}
