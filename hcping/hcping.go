package main

import (
	"net/http"
	"time"
)

func service1Handler(rw http.ResponseWriter, req *http.Request) {
	time.Sleep(3000)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hcping", service1Handler)
	http.ListenAndServe(":4000", mux)
}


