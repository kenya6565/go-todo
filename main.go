package main

import (
	"encoding/json"
	"go-todo/typefile"
	"html/template"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8090"

func start(w http.ResponseWriter, req *http.Request) {
	// JSON ファイルを読み出し用にオープン
	file, err := os.Open("article.json")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	var companyObj typefile.Company
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

	http.HandleFunc("/", start)
	log.Printf("connect to http://localhost:%s/", defaultPort)
	http.ListenAndServe(":8090", nil)
}
