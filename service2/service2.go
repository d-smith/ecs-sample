package main

import (
	"net/http"
)

func service2Handler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("service 2 response\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/service2", service2Handler)
	http.ListenAndServe(":5000", mux)
}


