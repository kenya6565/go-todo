package main

import (
	"html/template"
	"log"
	"net/http"
)

const defaultPort = "8090"

func hello(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func main() {

	http.HandleFunc("/", hello)
	log.Printf("connect to http://localhost:%s/", defaultPort)
	http.ListenAndServe(":8090", nil)
}
