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
	"strings"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	for {
		var router = mux.NewRouter()
		router.HandleFunc("/api/v1/status", statusCheck).Methods("GET")
		router.HandleFunc("/api/v1/start/{specs}", runContainer).Methods("GET")
		router.HandleFunc("/api/v1/remove/{id}", removeContainer).Methods("GET")
		router.HandleFunc("/api/v1/register", addKey).Methods("GET")

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

type Builder struct {
	specs   url.Values
	mux     sync.Mutex
	waiting bool
}

func (b *Builder) buildImage(path string, repoId string) {

	// Parsing ../../builds/ubuntu/tmp to ../../builds/ubuntu
	// Because we need to send the context dir to `docker build`
	pathSplit := strings.Split(path, "/")
	pathShort := pathSplit[:len(pathSplit)-1]
	pathContext := strings.Join(pathShort, "/")

	log.Printf("EXEC: docker build -t [%s] -f [%s] [%s]\n", repoId, path, pathContext)
	cmd := exec.Command("docker", "build", "-t", repoId, "-f", path, pathContext)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	log.Printf("Waiting for build to finish...\n")
	b.waiting = true
	go func() {
		for b.waiting {
			log.Printf("%v is building...\n", b.specs["name"])
			time.Sleep(time.Second * 3)
		}
	}()
	err = cmd.Wait()
	b.waiting = false
	log.Printf("Command finished with error: %v\n", err)
}

func runContainer(w http.ResponseWriter, r *http.Request) {

	// Parse the query string to create the specs key/values}
	query := mux.Vars(r)["specs"]
	specs, err := url.ParseQuery(query)
	if err != nil {
		panic(err)
	}

	builder := Builder{specs: specs}
	base := string(specs["base"][0])

	// If user entered a github repo...
	var image string
	if specs["id"] != nil {

		builder.mux.Lock()
		defer builder.mux.Unlock()
		repoId := string(specs["id"][0])
		image = repoId
		cloneURL := specs["cloneURL"][0]
		path := makeDockerfile(base, cloneURL, repoId)
		builder.buildImage(path, repoId)
		deadImage = image // deadImage is located in files.go. needed for clean

	} else {

		// Use the default image
		image = string(specs["base"][0])

	}

	// Execute docker run ...
	log.Printf("EXEC: docker run -itdP [%s]\n", image)
	runOut, runErr := exec.Command("docker", "run", "-itdP", image).Output()
	if runErr != nil {
		panic(runErr)
	}

	// Capture the hash and trim ofs trailing new line from output
	containerId := strings.TrimSpace(string(runOut))

	// Get the port number of the container's hash
	log.Printf("EXEC: docker port [%s]\n", containerId)
	portOut, portErr := exec.Command("docker", "port", containerId).Output()
	if portErr != nil {
		panic(portErr)
	}
	portNumberBytes := portOut[len(portOut)-6:]
	portNumber := strings.TrimSpace(string(portNumberBytes))

	// Representation of the running container
	type Container struct {
		Id         string
		Port       string
		Name       string
		Dockerfile string
	}
	container := Container{Id: containerId, Port: portNumber, Name: image, Dockerfile: image}

	// Send container data back to the frontend
	json.NewEncoder(w).Encode(container)
}

func removeContainer(w http.ResponseWriter, r *http.Request) {

	// Parse the query string to create the specs key/values
	query := mux.Vars(r)["id"]
	id := string(query)

	// Remove the container
	log.Printf("EXEC: docker rm -f %s\n", id)
	_, err := exec.Command("docker", "rm", "-f", id).Output()
	if err != nil {
		panic(err)
	}

	// Remove the temporary image if there is one
	if deadImage != "" {
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

func statusCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
}

type User struct {
	username string   `json:username`
	key      []string `json:keys`
}

// Gets user's private key from HTTP header and adds it to authorized_keys
func addKey(w http.ResponseWriter, r *http.Request) {

	key := strings.Replace(r.Header["Creds"][0], "\"[]", "", -1)
	key = strings.Replace(key, "\"", "", -1)
	key = strings.Replace(key, ",", "\n", -1)
	key = strings.Trim(key, "[]")
	key = strings.Trim(key, "\"")

	path := "/root/.ssh/authorized_keys"
	authorizedKeys, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open %s\n", path)
	}

	var keyString bytes.Buffer
	keyString.WriteString(key + "\n")
	_, writeErr := authorizedKeys.Write(keyString.Bytes())
	if writeErr != nil {
		log.Printf("Failed to write %s to %s\n", keyString.Bytes(), path)
		log.Print(err)
	}

	log.Printf("Added key to %s\n", path)

	json.NewEncoder(w).Encode("OK")
}
