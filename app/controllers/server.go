package controllers

import (
	"go-todo/config"
	"log"
	"net/http"
)

const defaultPort = "8090"

func StartMainServer() error {
	http.HandleFunc("/", top)
	log.Printf("connect to http://localhost:%s/", defaultPort)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
