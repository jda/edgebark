package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	fmt.Printf("listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
}

// accept it, parse it, "go" to fork handler and return right away.
// handlers for log, pagerduty, hipchat
// stage 2, web interface, logging, rules
