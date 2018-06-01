package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	//"os/exec"
)

var deadImage string

func makeDockerfile(base string, cloneURL string, repoId string) string {

	cloneURL = strings.Replace(cloneURL, "|", "/", -1)
	dockerFile := "../../builds/" + base + "/Dockerfile"
	tmpDockerfile := "../../builds/" + base + "/" + repoId

	// Read the original Dockerfile into a buffer
	content, err := ioutil.ReadFile(dockerFile)
	if err != nil {
		panic(err)
	}

	// Write the contents of Dockerfile to tmpDockerfile
	err = ioutil.WriteFile(tmpDockerfile, content, 0644)
	if err != nil {
		panic(err)
	}

	// Open tmp file and append RUN git clone ...
	tmp, err := os.OpenFile(tmpDockerfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	var gitClone bytes.Buffer
	gitClone.WriteString("RUN git clone " + cloneURL + "\n")
	_, writeErr := tmp.Write(gitClone.Bytes())
	if writeErr != nil {
		panic(writeErr)
	}

	log.Printf("Dockerfile created: [%s]\n", tmpDockerfile)
	deadImage = repoId
	return tmpDockerfile // "../../builds/ubuntu:14.04/1332132"
}
