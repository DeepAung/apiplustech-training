package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler)
	mux.HandleFunc("POST /", handler)
	mux.HandleFunc("PUT /{id}", handler)
	mux.HandleFunc("PATCH /{id}", handler)
	mux.HandleFunc("DELETE /{id}", handler)

	http.ListenAndServe(":8000", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	queries := []string{}
	for key, value := range r.URL.Query() {
		queries = append(queries, key+"="+value[0])
	}

	headers := []string{}
	for key, value := range r.Header {
		headers = append(headers, key+"="+value[0])
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	r.Body.Close()

	resString := fmt.Sprintf(
		"%s request\nurl: %s\nheaders: %s\nqueries: %s\nbody: %s\n",
		r.Method,
		r.URL.String(),
		headers,
		queries,
		string(reqBody),
	)

	w.Write([]byte(resString))
	return
}
