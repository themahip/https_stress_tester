package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	result "stess_tester/Result"
	Types "stess_tester/type"
	"sync"
	"time"
)

func POSTMETHOD(request *Types.Request) {
	wg := &sync.WaitGroup{}
	perresult := make(chan Types.PerResult, request.RequestPerUser*request.ConcurrentRequest)

	for i := 0; i < request.ConcurrentRequest; i++ {
		wg.Add(1)
		go postrequest(request, wg, perresult)
	}
	go func() {
		wg.Wait()
		close(perresult)
	}()

	result.Processresult(perresult, request)
}

func postrequest(request *Types.Request, wg *sync.WaitGroup, perresult chan<- Types.PerResult) {
	defer wg.Done()

	for i := 0; i < request.RequestPerUser; i++ {
		startTime := time.Now()
		payloadBytes, err := json.Marshal(request.Payload)
		if err != nil {
			fmt.Println("hyaa feri err")
		}
		payload := bytes.NewBuffer([]byte(payloadBytes))

		fmt.Println(payload)

		req, err := http.NewRequest("POST", request.Url, payload)
		if err != nil {
			fmt.Println("Error in creating new request:", err)
			perresult <- Types.PerResult{
				Err:        err,
				StatusCode: 0,
				Duration:   0,
			}
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Token "+request.AuthToken)

		client := http.Client{Timeout: time.Second * 20}
		res, err := client.Do(req)
		duration := time.Since(startTime)

		if err != nil {
			perresult <- Types.PerResult{
				Err:        err,
				StatusCode: 0,
				Duration:   duration,
			}
			fmt.Println("Error during HTTP request:", err)
		} else {

			defer res.Body.Close()

			perresult <- Types.PerResult{
				Err:        nil,
				StatusCode: res.StatusCode,
				Duration:   duration,
			}
		}
	}
}
