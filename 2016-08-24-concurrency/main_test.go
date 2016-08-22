package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// A test has a name and contains a series of requests to perform against the
// server.
type test struct {
	name     string
	requests []testRequest
}

// Each request in the test contains a path and a callback method to check the
// response.
type testRequest struct {
	method  string
	path    string
	payload io.Reader
	check   func(*http.Response) error
}

type testResult struct {
	name     string
	err      error
	duration time.Duration
}

// Starts a test server and executes the tests in serverTests
func TestServer(t *testing.T) {
	mux := createMux()

	server := httptest.NewServer(mux)
	t.Log("Server listening at:", server.URL)
	defer server.Close()

	tests := getTests()
	results := make(chan testResult)
	done := make(chan bool)
	numWorkers := 8

	// start a few worker go-routines
	for i := 0; i < numWorkers; i++ {
		go startWorker(server.URL, tests, results, done)
	}

	// close the results channel when all workers finish
	go func() {
		for i := 0; i < numWorkers; i++ {
			<-done
		}
		close(results)
	}()

	// collect all the results
	for result := range results {
		if result.err != nil {
			t.Fatalf("not ok %s (%dms): %s", result.name, result.duration, result.err)
		} else {
			t.Logf("ok %s (%dms)", result.name, result.duration)
		}
	}
}

// In our real codebase this would return a lot of tests, but here we'll just
// stub out 10 tests with a single request each.
func getTests() chan test {
	count := 10
	tests := make(chan test, count)

	for i := 0; i < count; i++ {
		tests <- test{
			name: fmt.Sprintf("Test %d", i+1),
			requests: []testRequest{{
				path: fmt.Sprintf("/test/%d", i+1),
				check: func(r *http.Response) error {
					buf, err := ioutil.ReadAll(r.Body)
					if err != nil {
						return err
					}
					if string(buf) != "Hello, WWG\n" {
						return fmt.Errorf(`Expected response body to be "Hello, WWG\n"`)
					}
					return nil
				},
			}},
		}
	}

	// signal that no more tests will be sent over the channel
	close(tests)

	return tests
}

func startWorker(serverURL string, tests chan test, results chan testResult, done chan bool) {
	for test := range tests {
		results <- runTest(serverURL, test)
	}
	done <- true
}

func runTest(serverURL string, test test) testResult {
	result := testResult{name: test.name}
	start := time.Now()

	for i, request := range test.requests {
		res, err := http.Get(serverURL + request.path)
		if err == nil {
			err = request.check(res)
		}
		if err != nil {
			result.err = fmt.Errorf("Request %d: %s", i, err)
		}
	}

	result.duration = time.Now().Sub(start) / 1000000
	return result
}
