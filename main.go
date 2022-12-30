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
	// u := &model.User{Name: "test", Email: "test@example.com", PassWord: "test@example.com"}
	// fmt.Println(u)
	// models.CreateUser(u)
	u, _ := models.GetUser(1)
	fmt.Println(u)
	u.Name = "test2"
	u.Email = "test2@example.com"
	models.UpdateUser(&u)
	u, _ = models.GetUser(1)
	fmt.Println(u)
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
	// fmt.Println(models.Db)
	// log.Printf("connect to http://localhost:%s/", defaultPort)
	http.ListenAndServe(":8090", nil)
}
