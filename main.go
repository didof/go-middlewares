package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func mwLogPath() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path)
			next(w, r)
		}
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello there")
}

func main() {
	r := mux.NewRouter()

	// create manually
	// mws := Middlewares{middlewares: []Middleware{mwLogPath()}}

	// or create with utility function
	mws := createMiddlewares(mwLogPath())

	// make and append middleware yourself
	mwFirst := makeMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am the first middleware")
	})
	mws.middlewares = append(mws.middlewares, mwFirst)

	// use utility function to declare and append it
	mws.addMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am the second middleware")
	})

	r.HandleFunc("/test/", mws.wrap(HelloHandler))

	log.Fatal(http.ListenAndServe(":8080", r))
}
