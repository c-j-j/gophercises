package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(writer, "Default Page")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	
	pathsToUrls := map[string]string {
		"/foo": "https://godoc.org/github.com/gophercises/urlshort",
	}

	mapHandler := MapHandler(pathsToUrls, mux)

	http.ListenAndServe(":8080", mapHandler)
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if foundUrl, ok := pathsToUrls[request.URL.Path]; ok {
			http.Redirect(writer, request, foundUrl, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(writer, request)
		}

	}
}