package main

import (
	"fmt"
	"net/http"
)

type sampleHandler int

func (m sampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test code inside this function")
}

func main() {
	var sh sampleHandler
	http.ListenAndServe(":8080", sh)
}
