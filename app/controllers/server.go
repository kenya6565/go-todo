package controllers

import (
	"fmt"
	"go-todo/app/models"
	"go-todo/config"
	"go-todo/model"
	"log"
	"net/http"
	"text/template"
)

const defaultPort = "8090"

// load HTML files and display them
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// create a session judging by a  saved cookie
func session(w http.ResponseWriter, r *http.Request) (session model.Session, err error) {
	// get a saved cookie
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		session = model.Session{UUID: cookie.Value}
		if ok, _ := models.CheckSession(&session); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return session, err

}

func StartMainServer() error {
	// load static files, which are bootstrap and jquery
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	// executed when posting login form
	http.HandleFunc("/authenticate", authenticate)

	log.Printf("connect to http://localhost:%s/", defaultPort)
	return http.ListenAndServe(":"+defaultPort, nil)
}
