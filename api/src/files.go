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

func makeDockerfile(base string, cloneURL string, repoName string) string {

  cloneURL = strings.Replace(cloneURL, "|", "/", -1)
  dockerFile := "../../builds/" + base + "/Dockerfile"


  // Use random string instead of "tmp" -- use Name from data
  tmpDockerfile := "../../builds/" + base + "/" + repoName

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

  fmt.Printf("TMP DOCKERFILE CREATED: [%s]\n", tmpDockerfile)
  return tmpDockerfile
}

