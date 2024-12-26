package methods

import (
	"fmt"
	"net/http"
	result "stess_tester/Result"
	Types "stess_tester/type"
	"sync"
	"time"
)

func GetMethod(Request *Types.Request) {
	wg := &sync.WaitGroup{}
	fmt.Println(Request)
	perresult := make(chan Types.PerResult, Request.ConcurrentRequest*Request.RequestPerUser)

	for i := 0; i < Request.ConcurrentRequest; i++ {
		wg.Add(1)
		go getrequest(perresult, Request, wg)
	}

	go func() {
		wg.Wait()
		close(perresult)
	}()
	result.Processresult(perresult, Request)
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
		req.Header.Set("Authorization", "Bearer "+request.AuthToken)

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
