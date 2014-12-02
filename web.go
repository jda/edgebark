package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/event/access/{type}", accessEvent)
	r.HandleFunc("/event/alarm/{type}", alarmEvent)
	r.HandleFunc("/event/door/{type}", doorEvent)
	r.HandleFunc("/event/system/{type}", systemEvent)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	fmt.Printf("listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// accept it, parse it, "go" to fork handler and return right away.
// handlers for log, pagerduty, hipchat
// stage 2, web interface, logging, rules
