package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		GiveResponseFile("responses/GET_alive.txt",w)
	})

	// List Sales
	http.HandleFunc("/api/sales", func(w http.ResponseWriter, r *http.Request){
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "405 method not allowed")
			return
		}

		if r.Header.Get("AuthToken") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "401 unauthorized")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		GiveResponseFile("responses/GET_api_sales.json",w)
	})

	// Facebook Login
	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request){
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "405 method not allowed")
			return
		}

		if r.Header.Get("FacebookToken") == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "400 bad request")
			return
		}

		if r.Header.Get("FacebookToken") == "bad_token" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "401 unauthorized")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		GiveResponseFile("responses/POST_api_login.json",w)
	})

	http.Handle("/", http.FileServer(http.Dir("../ui-frontend/src")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GiveResponseFile(filename string, w http.ResponseWriter){
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("reading sample file (%s)failed", filename))
	}
	fmt.Fprintf(w, string(file))
}