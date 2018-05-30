package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

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

		fmt.Println("Server running on port 5001...")
		log.Fatal(http.ListenAndServe(":5001",
			handlers.CORS(originsOk, headersOk, methodsOk)(router)),
			handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	}
}

func removeContainer(w http.ResponseWriter, r *http.Request) {

	// Parse the query string to create the specs key/values
	query := mux.Vars(r)["id"]
  id := string(query)

	// Remove the container
	fmt.Printf("EXEC: docker rm -f %s\n", id)
	_, err := exec.Command("docker", "rm", "-f", id).Output()
	if err != nil {
		panic(err)
	}

	// Remove the temporary image if there is one
  if deadImage != "" {
    fmt.Printf("EXEC: docker image rm -f %s\n", deadImage)
    out, err := exec.Command("docker", "image", "rm", "-f", deadImage).Output()
    if err != nil {
      panic(err)
    }
    fmt.Printf("image rm out: %s\n", out)
  }
	json.NewEncoder(w).Encode("OK")
}

func runContainer(w http.ResponseWriter, r *http.Request) {

	// Parse the query string to create the specs key/values
	query := mux.Vars(r)["specs"]
	specs, err := url.ParseQuery(query)
	if err != nil {
		panic(err)
	}
  base := string(specs["base"][0])

	// If user entered a github repo...
  var imageName string
	if specs["cloneURL"] != nil {

    repoName := string(specs["name"][0])
    cloneURL := specs["cloneURL"][0]
    path := makeDockerfile(base, cloneURL)
    imageName = buildImage(path, repoName)
    deadImage = imageName

  } else {

    // Use the default image
    imageName = string(specs["base"][0])

  }

	// Execute docker run ...
	fmt.Printf("EXEC: docker run -itdP %s\n", imageName)
	runOut, runErr := exec.Command("docker", "run", "-itdP", imageName).Output()
	if runErr != nil {
		panic(runErr)
	}

	// Capture the hash and trim ofs trailing new line from output
	containerId := strings.TrimSpace(string(runOut))

	// Get the port number of the container's hash
	fmt.Printf("EXEC: docker port [%s]\n", containerId)
	portOut, portErr := exec.Command("docker", "port", containerId).Output()
	if portErr != nil {
		panic(portErr)
	}
	portNumberBytes := portOut[len(portOut)-6:]
	portNumber := strings.TrimSpace(string(portNumberBytes))

	type Container struct {
		Id   string
		Port string
    Name string
	}
  container := Container{Id: containerId, Port: portNumber, Name: imageName}

	// Send container data back to the frontend
	json.NewEncoder(w).Encode(container)
}

func statusCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
	log.Print(w)
}
