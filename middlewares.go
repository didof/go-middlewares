package main

import (
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Middlewares struct {
	middlewares []Middleware
}

func createMiddlewares(mws ...Middleware) Middlewares {
	output := Middlewares{
		middlewares: mws,
	}
	return output
}

func (mws Middlewares) wrap(f http.HandlerFunc) http.HandlerFunc {
	for _, mw := range mws.middlewares {
		f = mw(f)
	}
	return f
}

func (mws *Middlewares) addMiddleware(f func(w http.ResponseWriter, r *http.Request)) {
	mw := makeMiddleware(f)
	mws.middlewares = append(mws.middlewares, mw)
}

func makeMiddleware(f func(w http.ResponseWriter, r *http.Request)) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f(w, r)
			next(w, r)
		}
	}
}

// MIDDLEWARES

func mwLogPath() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path)
			next(w, r)
		}
	}
}
