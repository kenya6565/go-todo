package controllers

import (
	"fmt"
	"go-todo/app/models"
	"go-todo/config"
	"go-todo/model"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

// create a session judging by a saved cookie
func session(w http.ResponseWriter, r *http.Request) (session model.Session, err error) {
	// get a saved cookie
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		log.Println(err)
		session = model.Session{UUID: cookie.Value}
		if ok, _ := models.CheckSession(&session); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return session, err
}

// start with either edit or update
// end with more than one integer
var validPath = regexp.MustCompile("^/todos/(edit|update)/([0-9]+)")

// type http.HandlerFunc returns func(ResponseWriter, *Request)
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// r.URL.Path = "/todos/edit/1"
		q := validPath.FindStringSubmatch(r.URL.Path)

		if q == nil {
			http.NotFound(w, r)
			return
		}

		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
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

	http.HandleFunc("/logout", logout)

	// only the user who can login can access
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)

	http.HandleFunc("/todos/edit/", parseURL(todoEdit))

	log.Printf("connect to http://localhost:%s/", defaultPort)
	return http.ListenAndServe(":"+defaultPort, nil)
}
