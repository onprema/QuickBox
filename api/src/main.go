package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
  "strings"
  "os"
  "net/url"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	for {
		var router = mux.NewRouter()
		router.HandleFunc("/api/v1/status", statusCheck).Methods("GET")
		router.HandleFunc("/api/v1/start/{specs}", runContainer).Methods("GET")

		// For CORS
		headersOk := handlers.AllowedHeaders([]string{"Authorization"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

		fmt.Println("Server running on port 5001...")
    log.Fatal(http.ListenAndServe(":5001",
			handlers.CORS(originsOk, headersOk, methodsOk)(router)),
      handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	}
}

func runContainer(w http.ResponseWriter, r *http.Request) {
  specs := mux.Vars(r)["specs"]
  m, err := url.ParseQuery(specs)
  if err != nil {
    panic(err)
  }
  image := m["base"]

  // Execute docker run
  runOut, runErr := exec.Command("docker", "run", "-itd", "-P", image[0]).Output()
	if runErr != nil {
		panic(runErr)
	}

  // Capture the hash and trim of trailing new line from output.
  imageId := strings.TrimSpace(string(runOut))

  // Get the port number
  portOut, portErr := exec.Command("docker", "port", imageId).Output()
  if portErr != nil {
    panic(portErr)
  }

  portNumberBytes := portOut[len(portOut)-6:]
  portNumber := strings.TrimSpace(string(portNumberBytes))
  json.NewEncoder(w).Encode(portNumber)
}

type Response struct {
  Name string
  Port uint
}

func statusCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
  log.Print(w)
}
