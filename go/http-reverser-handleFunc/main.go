package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /reverser", ArgValidator(reverser))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func reverser(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("arg")
	rv := reverse(a)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(rv))
}

func ArgValidator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a := r.URL.Query().Get("arg")
		if a == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("arg is required"))
			return
		}
		next.ServeHTTP(w, r)
	}
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
