package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Entrypoint. Sets up routes and runs the daemon that reaps containers.
func main() {
	for {
		var router = mux.NewRouter()
		router.HandleFunc("/api/v1/status", statusCheck).Methods("GET")
		router.HandleFunc("/api/v1/start/{specs}", runContainer).Methods("GET")
		router.HandleFunc("/api/v1/remove/{id}", removeContainer).Methods("GET")

		// For CORS
		headersOk := handlers.AllowedHeaders([]string{"Authorization"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

		// Cleanup function for containers > 12 hours old
		go func() {
			for {
				dead := killContainers()
				log.Printf("Killed %v containers\n", dead)
				time.Sleep(time.Minute * 10)
			}
		}()

		fmt.Println("Server running on port 5001...")
		log.Fatal(http.ListenAndServe(":5001",
			handlers.CORS(originsOk, headersOk, methodsOk)(router)),
			handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	}
}

// Check the status of the app
func statusCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
}
