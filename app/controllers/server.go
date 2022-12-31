package controllers

import (
	"go-todo/config"
	"log"
	"net/http"
)

const defaultPort = "8090"

func StartMainServer() error {
	// load static files
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)

	log.Printf("connect to http://localhost:%s/", defaultPort)
	return http.ListenAndServe(":"+defaultPort, nil)
}
