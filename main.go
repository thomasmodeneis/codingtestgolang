package main

import (
	"net/http"
	"codingtestgolang/routes"
	"os"
	"log"
)

var Port = os.Getenv("PORT")

func NewServer(Port string) *http.Server {
	router := routes.Handler()
	return &http.Server{
		Addr:    ":"+Port,
		Handler: router,
	}
}

func init(){
	if Port == "" {
		Port = "3000"
	}
}

func main() {
	s := NewServer(Port)
	log.Println("Server starting --> " + Port)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln("Error: %v", err)
		os.Exit(0)
	}
}

