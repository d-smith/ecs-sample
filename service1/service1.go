package main

import (
	"net/http"
)

func service1Handler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("service 1 response\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/service1", service1Handler)
	http.ListenAndServe(":4000", mux)
}


