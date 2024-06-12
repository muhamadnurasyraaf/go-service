package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func jsonResponse(w http.ResponseWriter, message string) {
	response := map[string]string{"message": message}
	jsonData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

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
		jsonResponse(w, "Hello World")
	})

	color.Cyan("***********************************")
	color.Cyan(" Starting server at %s", addr)
	color.Cyan("***********************************")

	// Start the server
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
