package main

import (
	"encoding/json"
	"go-todo/model"
	"html/template"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8090"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func start(w http.ResponseWriter, req *http.Request) {
	// JSON ファイルを読み出し用にオープン
	file, err := os.Open("article.json")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	var companyObj []model.Company
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&companyObj); err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", companyObj)

	if err := tpl.Execute(w, companyObj); err != nil {
		panic(err.Error())
	}
}

func main() {
	http.HandleFunc("/", start)
	log.Printf("connect to http://localhost:%s/", defaultPort)
	http.ListenAndServe(":8090", nil)
}
