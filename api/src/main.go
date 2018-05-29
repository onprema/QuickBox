package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
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
	// Gets container ID from the http Handler and removes the container

	// Parse the query string to create the specs key/values
	query := mux.Vars(r)["id"]
  /*
	specs, err := url.ParseQuery(query)
	if err != nil {
		panic(err)
	}
  */

  id := string(query)

	// Remove the container
	fmt.Printf("EXEC: docker rm -f %s\n", id)
	_, err := exec.Command("docker", "rm", "-f", string(id)).Output()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode("OK")
}

func buildImage(user string, repo string) string {
	// Takes the github user/repo and returns the name of the image to
  // runContainer

	// Create buffers to read bytes and convert to strings
	var tagBuffer bytes.Buffer
	var imageBuffer bytes.Buffer

	// Build the tag string
	for i := 0; i < len(user); i++ {
		tagBuffer.WriteByte(user[i])
	}
	tagBuffer.WriteString("/")
	for i := 0; i < len(repo); i++ {
		tagBuffer.WriteByte(repo[i])
	}
	// Image tag must be lowercase
	tag := tagBuffer.String()
	tag = strings.ToLower(tag)

	// Build the image name string (default to :latest)
	for i := 0; i < len(tag); i++ {
		imageBuffer.WriteByte(tag[i])
	}
	imageBuffer.WriteString(":latest")
	imageName := imageBuffer.String()

	// Define the path to the Dockerfile of the base image
	imagePath, _ := filepath.Abs("../../builds/ubuntu/14.04")

	// Run the command
	fmt.Printf("EXEC: docker build -t %s %s\n", tag, imagePath)
	cmd := exec.Command("docker", "build", "-t", tag, imagePath)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	// Return the image name to runContainer() (ex: "my_name/my_repo:latest")
	return imageName
}

func runContainer(w http.ResponseWriter, r *http.Request) {
	// runContainer gets specifications from the http Handler and runs a
	// a container based on those specs

	// Parse the query string to create the specs key/values
	query := mux.Vars(r)["specs"]
	specs, err := url.ParseQuery(query)
	if err != nil {
		panic(err)
	}

	// If user entered a github repo...
	var imageName string
	if specs["repo"] != nil {

		// repoUser and repoName has been split, like (my_name:my_repo)
		repoUser := string(specs["repo"][0])
		repoUser = strings.Split(repoUser, ":")[0]
		repoName := string(specs["repo"][0])
		repoName = strings.Split(repoName, ":")[1]
		repoName = strings.ToLower(repoName)

		// Call buildImage to create an image named "my_name/my_repo:latest"
		// The container should include the cloned github repo
		imageName = buildImage(repoUser, repoName)
	} else {
		imageName = string(specs["base"][0])
	}

	// Execute docker run ...
	fmt.Printf("EXEC: docker run -itdP %s\n", imageName)
	runOut, runErr := exec.Command("docker", "run", "-itdP", imageName).Output()
	if runErr != nil {
		panic(runErr)
	}

	// Capture the hash and trim of trailing new line from output
	containerId := strings.TrimSpace(string(runOut))

	// Get the port number of the container's hash
	fmt.Printf("EXEC: docker port [%s]\n", containerId)
	portOut, portErr := exec.Command("docker", "port", containerId).Output()
	if portErr != nil {
		panic(portErr)
	}

	// Get the last 5 characters of the output (the port number)
	portNumberBytes := portOut[len(portOut)-6:]
	portNumber := strings.TrimSpace(string(portNumberBytes))

	type Container struct {
		Id   string
		Port string
	}
	container := Container{Id: containerId, Port: portNumber}

	// Send container data back to the frontend
	json.NewEncoder(w).Encode(container)
}

func statusCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
	log.Print(w)
}
