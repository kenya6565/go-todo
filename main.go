package main

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

const defaultPort = "8090"
const message = "This is a test message"

func hello(w http.ResponseWriter, req *http.Request) {
	p1 := Person{
		Name: "hogefuga",
		Age:  28,
	}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, p1); err != nil {
		panic(err.Error())
	}
}

func main() {

	http.HandleFunc("/", hello)
	log.Printf("connect to http://localhost:%s/", defaultPort)
	http.ListenAndServe(":8090", nil)
}
