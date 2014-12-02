package main

import (
	"fmt"
	"net/http"
)

func systemEvent(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "systemHandler")
}
