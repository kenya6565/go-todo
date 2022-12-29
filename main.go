package main

import (
	"encoding/json"
	"fmt"
	"go-todo/app/models"
	"go-todo/config"
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

func convertJson(file *os.File) []model.Company {
	var companyObj []model.Company
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&companyObj); err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", companyObj)
	return companyObj
}

func start(w http.ResponseWriter, req *http.Request) {
	// Open Json file
	file, err := os.Open("article.json")
	if err != nil {
		panic(err.Error())
	}
	// Close at the end of start method
	defer file.Close()

	companyStruct := convertJson(file)
	if err := tpl.Execute(w, companyStruct); err != nil {
		panic(err.Error())
	}
}

func main() {
	http.HandleFunc("/", start)
	fmt.Println(config.Config.Port)
	fmt.Println(models.Db)
	log.Printf("connect to http://localhost:%s/", defaultPort)
	http.ListenAndServe(":8090", nil)
	u := &model.User()
	u.Name = "test"
	u.Email = "test@test.com"
	u.Password = "test"
}
