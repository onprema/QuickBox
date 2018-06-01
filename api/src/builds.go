package main

/*

import (
  //"fmt"
//  "os"
 // "bytes"
 // "io/ioutil"
  "strings"
  "os/exec"
 // "sync"
 // "net/http"
 "log"
)

func buildImage(path string, repoId string) {

  // Parsing ../../builds/ubuntu/tmp to ../../builds/ubuntu
  // Because we need to send the context dir to `docker build`
  pathSplit := strings.Split(path, "/")
  pathShort := pathSplit[:len(pathSplit) - 1]
  pathContext := strings.Join(pathShort, "/")

  log.Printf("EXEC: docker build -t [%s] -f [%s] [%s]\n", repoId, path, pathContext)
	cmd := exec.Command("docker", "build", "-t", repoId, "-f", path, pathContext)
  err := cmd.Start()
	if err != nil {
		panic(err)
	}
  log.Printf("Waiting for build to finish\n")
  err = cmd.Wait()
  log.Printf("Command finished with error: %v\n", err)
}

/*
type Build struct {
  mux sync.RWMutex
  cmd *exec.Cmd
  status BuildStatus
}

type Server struct {
  mux sync.RWMutex
  builds map[string]*Build
}

func (server *Server) jobStatus(w http.ResponseWriter, r *http.Request) {
  buildID := getBuildID(r)
  server.mux.RLock()
  build, err := server.builds[buildID]
  if err != nil {
    http.NotFound(r, w)
    return
  }
  status := build.Status()
  fmt.Printf("build: [%v]\nstatus: [%s]\n", build, status)
}

func (server *Server) createJob(w http.ResponseWriter, r *http.Request) {
  build := &Build{cmd: exec.Command("docker", "build", "-t", repoId, "-f", path, context)}
  server.mux.Lock()
  id := server.newBuildID()
  server.builds[id] = build
  server.mux.Unlock()
  go b.Run()
  w.Header().Set("Location", "/jobs/" + id)
  w.WriteHeader(http.StatusCreated)
}

func (server *Server) newBuildID() string {
  // Implement this. Only call it with server.mux held
  return "Some bogus Id"
}

func getBuildID(r *http.Request) string {
  // Implement this
  return "some bogus name"
}

func (b *Build) Run() {
  b.cmd.Start()
  b.mux.Lock()
  b.status = Running
  b.mux.Unlock()
  err := b.cmd.Wait()
  b.mux.Lock()
  if err != nil {
    b.status = Failed
  } else {
    b.status = Done
  }
  b.mux.Unlock()
}

func (b *Build) Status() BuildStatus {
  b.mux.RLock()
  defer b.mux.RUnlock()
  return b.status
}

type BuildStatus int

const (
  Running BuildStatus = iota
  Done
  Failed
)
*/
