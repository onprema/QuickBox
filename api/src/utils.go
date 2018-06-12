package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var deadImage string

// Creates a dockerfile for users who want to import a repository
func makeDockerfile(base string, cloneURL string, repoId string, pw string) string {

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

	// Add repo to dockerfile if they provided one
	if cloneURL != "tmp" {
		if strings.Contains(cloneURL, "\n") { return "err" }
		var gitClone bytes.Buffer
		gitClone.WriteString("RUN git clone " + cloneURL + " && cd /\n")
		_, cloneWriteErr := tmp.Write(gitClone.Bytes())
		if cloneWriteErr != nil {
			panic(cloneWriteErr)
		}
	}

	// Change root pw
	var rootPW bytes.Buffer
	rootPW.WriteString("RUN echo \"root:" + pw + "\" | chpasswd\n")
	_, pwWriteErr := tmp.Write(rootPW.Bytes())
	if pwWriteErr != nil {
		panic(pwWriteErr)
	}

	log.Printf("Dockerfile created: [%s]\n", tmpDockerfile)
	deadImage = repoId
	return tmpDockerfile // "../../builds/ubuntu:14.04/1332132"
}

// Kills any containers than have been running longer than 12 hours
func killContainers() int {
	output, err := exec.Command("docker", "ps", "-a", "--format", "\"{{.ID}} {{.RunningFor}}\"").Output()
	if err != nil {
		panic(err)
	}
	containers := strings.Split(strings.Trim(string(output), " \"\n"), "\n")
	numDead := 0
	if len(containers) > 1 {
		for i := 0; i < len(containers); i++ {
			timeString := strings.Split(strings.Trim(containers[i], "\" "), " ")
			containerID := timeString[0]
			if timeString[1] == "5" && timeString[2] == "minutes" {
				cmd := exec.Command("docker", "container", "rm", "-f", containerID)
				err := cmd.Run()
				if err != nil {
					panic(err)
				}
				numDead += 1
			}
		}
	}
	return numDead
}
