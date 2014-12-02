package main

import (
	"fmt"
	"net/http"
	"os"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/event/{id}/access/{type}", accessEvent)
  r.HandleFunc("/event/{id}/alarm/{type}", alarmEvent)
  r.HandleFunc("/event/{id}/door/{type}", doorEvent)
  r.HandleFunc("/event/{id}/system/{type}", systemEvent)
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
// only config env besides port should be rethinkdb url. settings in there.