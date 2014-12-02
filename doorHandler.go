package main

import (
	"fmt"
	"net/http"
)

func doorEvent(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "doorHandler")
}
