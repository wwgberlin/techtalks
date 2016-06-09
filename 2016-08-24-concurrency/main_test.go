package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
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

// Starts a test server and executes the tests in serverTests
func TestServer(t *testing.T) {
	mux := createMux()

	server := httptest.NewServer(mux)
	t.Log("Server listening at:", server.URL)
	defer server.Close()

	tests := getTests()

	for _, test := range tests {
		if err := runTest(server.URL, test); err != nil {
			t.Fatalf("not ok %s: %s", test.name, err)
		} else {
			t.Logf("ok %s", test.name)
		}
	}
}

// In our real codebase this would return a lot of tests, but here we'll just
// stub out 10 tests with a single request each.
func getTests() []test {
	count := 10
	tests := make([]test, count)

	for i := 0; i < count; i++ {
		tests[i] = test{
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

	return tests
}

func runTest(serverURL string, test test) error {
	for i, request := range test.requests {
		res, err := http.Get(serverURL + request.path)
		if err != nil {
			return fmt.Errorf("Request %d: %s", i, err)
		}
		err = request.check(res)
		if err != nil {
			return fmt.Errorf("Request %d: %s", i, err)
		}
	}
	return nil
}
