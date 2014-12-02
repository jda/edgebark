package main

import (
	"fmt"
	"net/http"
)

func accessEvent(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "accessHandler")
}
