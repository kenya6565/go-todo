package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Company struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

const defaultPort = "8090"

func hello(w http.ResponseWriter, req *http.Request) {
	// JSON ファイルを読み出し用にオープン
	file, err := os.Open("article.json")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	var companyObj Company
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&companyObj); err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", companyObj)

	t, err := template.ParseFiles("index.html")

	if err := t.Execute(w, companyObj); err != nil {
		panic(err.Error())
	}
}

func main() {

	http.HandleFunc("/", hello)
	log.Printf("connect to http://localhost:%s/", defaultPort)
	http.ListenAndServe(":8090", nil)
}
