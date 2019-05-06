package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func decorated(h http.HandlerFunc) http.Handler {
	// Decorate the original function with some added functionality
	// Wrap the original function with the logging facility and
	// once that's done call it.
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}
func protect(h http.HandleFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		// protected route code i.e. middleware
		// middleware only decorates the route.
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc('/', decorated(hello))
	server.ListenAndServe()
}
