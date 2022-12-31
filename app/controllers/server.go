package controllers

import (
	"log"
	"net/http"
)

const defaultPort = "8090"

func StartMainServer() error {
	http.HandleFunc("/", top)
	log.Printf("connect to http://localhost:%s/", defaultPort)
	return http.ListenAndServe(":"+defaultPort, nil)
}
