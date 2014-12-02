package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var cfg config

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/event/{id}/{category}/{type}", newEvent)
	http.Handle("/", r)

	cfg, err := loadConfig(cfg)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("listening on port %s...\n", cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Panic(err)
	}
}

// accept it, parse it, "go" to fork handler and return right away.
// handlers for log, pagerduty, hipchat
// stage 2, web interface, logging, rules
// only config env besides port should be rethinkdb url. settings in there.
