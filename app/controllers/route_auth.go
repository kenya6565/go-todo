package controllers

import (
	"fmt"
	"go-todo/app/models"
	"go-todo/model"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)

		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
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

// show login page
func login(w http.ResponseWriter, req *http.Request) {
	_, err := session(w, req)

	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, req, "/todos", 302)
	}
}

func logout(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := model.Session{UUID: cookie.Value}
		fmt.Println("session", session)
		models.DeleteSessionByUUID(&session)
	}
	http.Redirect(w, req, "/login", 302)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	inputted_email := r.PostFormValue("email")

	// create a user, using an inputted email from a login form
	user, err := models.GetUserByEmail(inputted_email)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}

	// save a user session in a cookie when the inputted password is correct
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := models.CreateSession(&user)
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
