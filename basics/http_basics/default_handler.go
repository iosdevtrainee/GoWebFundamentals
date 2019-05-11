package main

import (
	"fmt"
	"net/http"
)

type DefaultHandler struct{}

// with this method the method DefaultHandler implements the Handler interface

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// Source for http.HandleFunc

func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	mux.Handle(pattern, HandlerFunc(handler))
}

func main() {
	handler := DefaultHandler{}
	server := http.Server{
		Addr: "127.0.0.1",
		// the default value for Handler is ServeMux (DefaultServeMux)
		//  which also implements
		// the Handler interface however it's handler provides router functionality
		Handler: &handler,
	}
	server.ListenAndServe()
}
