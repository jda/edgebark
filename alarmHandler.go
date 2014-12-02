package main

import (
	"fmt"
	"net/http"
)

func alarmEvent(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "alarmHandler")
}
