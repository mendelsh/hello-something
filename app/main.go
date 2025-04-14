package main

import (
	"log"
	"net/http"
	"app/server"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})
	mux := http.NewServeMux()
	mux.HandleFunc("/create", server.CreateRoomHandler)
	mux.HandleFunc("/join", server.JoinRoomHandler)

	handler := c.Handler(mux)
	log.Println("Starting server on :8082")

	log.Fatal(http.ListenAndServe("0.0.0.0:8082", handler))
}