package controllers

import (
	"encoding/json"
	"go-todo/model"
	"log"
	"net/http"
	"os"
	"text/template"
)

var tpl *template.Template

func convertJson(file *os.File) []model.Company {
	var companyObj []model.Company
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&companyObj); err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", companyObj)
	return companyObj
}

// pass html elements generated from .json file
func top(w http.ResponseWriter, req *http.Request) {
	// Open Json file
	file, err := os.Open("article.json")
	if err != nil {
		panic(err.Error())
	}
	// Close at the end of start method
	defer file.Close()

	companyStruct := convertJson(file)
	// if err := tpl.Execute(w, companyStruct); err != nil {
	// 	panic(err.Error())
	// }
	generateHTML(w, companyStruct, "layout", "top")
}
