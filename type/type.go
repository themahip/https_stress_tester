package Types

import "time"

type Request struct {
	Url               string
	AuthToken         string
	Method            string
	Payload           map[string]interface{}
	RequestPerUser    int
	ConcurrentRequest int
}

type PerResult struct {
	Err        error
	StatusCode int
	Duration   time.Duration
}

type FinalResult struct {
	TargetUrl             string
	TotalRequestSent      int
	ConcurrentRequestSent int
	FailedRequest         int
	SuccessRequest        int
	TotalDuration         time.Duration
}
