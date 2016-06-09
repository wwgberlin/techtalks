package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

// In our real application this loads configuration, adds routes and so on. Here
// we will just include a trivial handler that sleeps to simulate work and
// responds with "Hello, WWG"
func createMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		time.Sleep(
			time.Duration(rand.Float32() * 3 * float32(time.Second)),
		)

		w.Write([]byte("Hello, WWG\n"))
	})

	return mux
}

func main() {
	mux := createMux()
	log.Fatal(http.ListenAndServe(":8001", mux))
}
