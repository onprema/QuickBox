package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

// Representation of the running container
type Container struct {
	Id         string
	Port       string
	Name       string
	Dockerfile string
	Password   string
}

// Builds and runs a container
func runContainer(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["specs"]
	specs, err := url.ParseQuery(query)
	if err != nil {
		panic(err)
	}

	builder := Builder{specs: specs}
	base := string(specs["base"][0])
	pw := string(specs["pw"][0])
	log.Printf("PW: [%s]\n", pw)

	var cloneURL string
	var repoId string
	var image string
	if specs["id"] != nil {
		repoId = string(specs["id"][0])
		image = repoId
		cloneURL = specs["cloneURL"][0]
	} else {
		image = string(specs["base"][0])
		repoId = "tmp"
		cloneURL = "tmp"
	}
	builder.mux.Lock()
	defer builder.mux.Unlock()
	path := makeDockerfile(base, cloneURL, repoId, pw)
	builder.buildImage(path, repoId)
	deadImage = repoId // deadImage is located in files.go.

	// Execute docker run, limiting memory and CPU
	runOut, runErr := exec.Command("docker", "run", "-itdP", "--memory", "10m", "--cpus", ".5", strings.Split(path, "/")[4]).Output()
	log.Printf("EXEC: docker run -itdP --memory 10m --cpus .5 [%s]\n", strings.Split(path, "/")[4])
	if runErr != nil {
		log.Fatal(runErr)
	}
	containerId := strings.TrimSpace(string(runOut))
	portOut, portErr := exec.Command("docker", "port", containerId).Output()
	log.Printf("EXEC: docker port [%s]\n", containerId)
	if portErr != nil {
		panic(portErr)
	}
	portNumberBytes := portOut[len(portOut)-6:]
	portNumber := strings.TrimSpace(string(portNumberBytes))

	container := Container{Id: containerId, Port: portNumber, Name: image, Dockerfile: image, Password: pw}
	json.NewEncoder(w).Encode(container)
}

// Kills a container (when user clicks 'Destroy' button)
func removeContainer(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["id"]
	id := string(query)

	_, err := exec.Command("docker", "rm", "-f", id).Output()
	log.Printf("EXEC: docker rm -f %s\n", id)
	if err != nil {
		panic(err)
	}
	// Remove the temporary image if there is one
	if deadImage != "" && deadImage != "tmp" {
		log.Printf("deadImage: [%s]\n", deadImage)
		log.Printf("EXEC: docker image rm -f %s\n", deadImage)
		_, err := exec.Command("docker", "image", "rm", "-f", deadImage).Output()
		if err != nil {
			panic(err)
		}
		deadImage = ""
	}
	json.NewEncoder(w).Encode("OK")
}
