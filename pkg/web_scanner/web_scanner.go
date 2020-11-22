package web_scanner

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

type RequestResult struct {
	RequestedURL *url.URL
	HTTPStatus   string
	ElapsedTime  time.Duration
	Err          error
}

type WebScanner struct {
	*sync.WaitGroup
	http.Client
	resultCh chan *RequestResult
}

func NewWebScanner(timeOut int) *WebScanner {
	return &WebScanner{
		WaitGroup: &sync.WaitGroup{},
		Client: http.Client{
			Timeout: time.Second * time.Duration(timeOut),
		},
		resultCh: make(chan *RequestResult),
	}
}

func (ws *WebScanner) MakeAllRequests(urls []*url.URL) []*RequestResult {
	wg := ws.WaitGroup
	for _, reqURL := range urls {
		wg.Add(1)
		go ws.makeRequestAndGetResults(wg, reqURL)
	}

	results := make([]*RequestResult, len(urls))
	i := 0
	for reqResult := range ws.resultCh {
		results[i] = reqResult
		i++
	}

	return results
}

func (ws *WebScanner) makeRequestAndGetResults(wg *sync.WaitGroup, reqURL *url.URL) {
	defer wg.Done()

	start := time.Now()
	status, err := ws.makeRequest(reqURL)
	end := time.Now().Sub(start)

	ws.resultCh <- &RequestResult{
		RequestedURL: reqURL,
		HTTPStatus:   status,
		ElapsedTime:  end,
		Err:          err,
	}
}

func (ws *WebScanner) makeRequest(reqURL *url.URL) (string, error) {
	resp, err := ws.Get(reqURL.String())
	if err != nil {
		return "", err
	}

	return resp.Status, nil
}
