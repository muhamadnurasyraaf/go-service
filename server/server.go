package server

import (
	"fmt"
	"full-webapp/api"
	"full-webapp/router"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func StartServer() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("SERVER_HOST")

	port := os.Getenv("SERVER_PORT")

	if host == "" {
		log.Fatal("SERVER_HOST is not set in the .env file")
	}

	if port == "" {
		log.Fatal("SERVER_PORT is not set in the .env file")
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.JsonParse(w, "Hello World")
	})

	color.Cyan("***********************************")
	color.Cyan(" Starting server at %s", addr)
	color.Cyan("***********************************")

	r := router.NewRouter()

	router.RegisterRoute(r)

	// Start the server
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
