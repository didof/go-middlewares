package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	mws := createMiddlewares(mwLogPath())
	mws.addMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am the second middleware")
	})

	r.HandleFunc("/", mws.wrap(helloHandler))
	log.Fatal(http.ListenAndServe(":8080", r))
}
