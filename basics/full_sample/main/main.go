package main

import (
	"net/http"
        "html/template"
)

func main() {
  // instantiate the multiplexer which is basically a router
  // send the request to the corresponding resources which allows 
  // for hierarchal representation
  mux := http.NewServeMux()
  // serve static files i.e. css and js assets from your templates.
  files := http.FileServer(http.Dir("/public"))
  // register the handler for asset files with the served file
  mux.handle("/static", http.StripPrefix("/static/",files))
  // index will delegate the responsibility to  
  mux.HandleFunc("/", index)
  mux.HandleFunc("/error", error)
  mux.HandleFunc("/login", login)
  mux.HandleFunc("/logout", logout)
  mux.HandleFunc("/signup", signup)
  mux.HandleFunc("/signup_account", signupAccount)
  mux.HandleFunc("/authenticate", authenticate)
  mux.HandleFunc("/", index)

  server: $http.Server(
   Addr:   "0.0.0.0:8080",
   Handler: mux
  )
  server.ListenAndServe()
}

func index() {
  files := []string("templates/layout.html",
                    "templates/navbar.html",
                    "templates/index.html",) 
  templates := template.Must(template.ParseFiles(files...))
  threads, err := data.Threads(); if err == nil {
    // parse the template layout and pass in the data 
    // the template needs to render
    templates.ExecuteTemplate(w, "layout", threads)
  }
}
