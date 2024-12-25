package method

import (
	"fmt"
	"net/http"
	Types "stess_tester/type"
	"sync"
	"time"
)

func GETMETHOD(Request *Types.Request) {
	wg := &sync.WaitGroup{}
	fmt.Println(Request)
	result := make(chan Types.PerResult, Request.ConcurrentRequest*Request.RequestPerUser)
	startTime := time.Now()
	for i := 0; i < Request.ConcurrentRequest; i++ {
		wg.Add(1)
		go getrequest(result, Request, wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	processresult(result, Request, startTime)

}

func getrequest(result chan<- Types.PerResult, request *Types.Request, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{Timeout: time.Second * 10}
	startTime := time.Now()
	for i := 0; i < request.RequestPerUser; i++ {
		req, err := http.NewRequest("GET", request.Url, nil)
		if err != nil {
			fmt.Println("error in creating new request")
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Token "+request.AuthToken)

		res, err := client.Do(req)
		duration := time.Since(startTime)
		if err != nil {
			result <- Types.PerResult{
				Err:        err,
				StatusCode: 0,
				Duration:   duration,
			}
		} else {
			result <- Types.PerResult{
				Err:        nil,
				StatusCode: res.StatusCode,
				Duration:   duration,
			}
		}

	}

}

func processresult(perresults <-chan Types.PerResult, request *Types.Request, starttime time.Time) {
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
	fmt.Println(finalResult)
}
