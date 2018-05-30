package main

import (
  "fmt"
  "os"
  "bytes"
  "io/ioutil"
  "strings"
  "os/exec"
)

var deadImage string;

func makeDockerfile(base string, cloneURL string) string {

  cloneURL = strings.Replace(cloneURL, "|", "/", -1)
  dockerFile := "../../builds/" + base + "/Dockerfile"
  tmpDockerfile := "../../builds/" + base + "/tmp"

  // Read the original Dockerfile into a buffer
  content, err := ioutil.ReadFile(dockerFile)
  if err != nil {
    panic(err)
  }

  // Write the contents of Dockerfile to tmp
  err = ioutil.WriteFile(tmpDockerfile, content, 0644)
  if err != nil {
    panic(err)
  }

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

  fmt.Printf("tmpDockerfile created.")
  return tmpDockerfile
}

func buildImage(path string, repoName string) string {

  // Parsing ../../builds/ubuntu/tmp to ../../builds/ubuntu
  pathSplit := strings.Split(path, "/")
  pathShort := pathSplit[:len(pathSplit) - 1]
  pathLess := strings.Join(pathShort, "/")

  fmt.Printf("EXEC: docker build -t %s -f %s %s\n", repoName, path, pathLess)
	buildOut, buildErr := exec.Command("docker", "build", "-t", repoName, "-f", path, pathLess).Output()
	if buildErr != nil {
		panic(buildErr)
	}
  fmt.Println(string(buildOut))

	// Return the image name to runContainer() (ex: "my_name/my_repo:latest")
  return repoName
}
