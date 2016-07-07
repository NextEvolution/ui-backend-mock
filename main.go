package main

import (
	"net/http"
	"fmt"
	"html"
	"log"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/api/sales", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json, err := ioutil.ReadFile("static_content/list_sales.json")
		if err != nil {
			panic("reading sample html file failed")
		}
		fmt.Fprintf(w, string(json))
	})

	http.Handle("/", http.FileServer(http.Dir("../ui-frontend/src")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}