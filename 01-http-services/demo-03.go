package main

import (
	"fmt"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func profileTime(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			end := time.Now()
			fmt.Printf("Request took %s\n", end.Sub(start))
		}()
		f(w, r)
	}
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request made for : %s\n", r.URL.Path)
		f(w, r)
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func main() {
	/*
		http.HandleFunc("/foo", profileTime(logging(foo)))
		http.HandleFunc("/bar", profileTime(logging(bar)))
	*/
	http.HandleFunc("/foo", Chain(foo, logging, profileTime))
	http.HandleFunc("/bar", Chain(bar, logging, profileTime))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "bar")
}
