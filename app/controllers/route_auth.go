package controllers

import (
	"go-todo/app/models"
	"go-todo/model"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user := model.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := models.CreateUser(&user); err != nil {
			log.Println(err) 
		}
	}
	http.Redirect(w, r, "/", 302)

}
