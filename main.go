package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"nextevolution/ui-backend-mock/types"
	"encoding/json"
	"os"
)

func main() {
	configPath := os.Args[1]
	if configPath == "" {
		log.Panic("Please supply a config file path")
	}

	rawConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panic(fmt.Sprintf("reading config file (%s)failed", configPath))
	}

	config := &types.Config{}
	err = json.Unmarshal(rawConfig, &config)
	if err != nil {
		log.Panic(fmt.Sprintf("unable to unmarshal config file (%s)", configPath))
	}

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

		if r.Header.Get("AuthToken") == "bad_token" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "401 unauthorized")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		GiveResponseFile("responses/GET_api_sales.json",w)
	})

	// Facebook Login
	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request){
		if r.Method != "POST" || r.Method != "OPTIONS"{
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "405 method not allowed")
			return
		}

		//unmarshal data
		var loginReq types.FbLoginReq

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "400 bad request - Can't read body")
			return
		}

		err = json.Unmarshal(body, &loginReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "400 bad request - Can't read json")
			return
		}

		if loginReq.FbToken == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "400 bad request - empty token")
			return
		}

		if loginReq.FbToken == "bad_token" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "401 unauthorized")
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		GiveResponseFile("responses/POST_api_login.json",w)
	})

	http.Handle("/", http.FileServer(http.Dir(config.StaticFilePath)))

	log.Printf("Listening on port: %d", config.Port)
	log.Printf("Static file path: %s", config.StaticFilePath)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), Log(http.DefaultServeMux)))
}

func GiveResponseFile(filename string, w http.ResponseWriter){
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("reading sample file (%s)failed", filename))
	}
	fmt.Fprintf(w, string(file))
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}