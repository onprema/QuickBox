package main

import (
	"log"
	"net/url"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// Data structure for building images
type Builder struct {
	specs   url.Values
	mux     sync.Mutex
	waiting bool
}

// Builds an image from a Dockerfile
func (b *Builder) buildImage(path string, repoId string) {
	pathSplit := strings.Split(path, "/")
	pathShort := pathSplit[:len(pathSplit)-1]
	pathContext := strings.Join(pathShort, "/")

	cmd := exec.Command("docker", "build", "-t", repoId, "-f", path, pathContext)
	log.Printf("EXEC: docker build -t [%s] -f [%s] [%s]\n", repoId, path, pathContext)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	b.waiting = true
	go func() {
		for b.waiting {
			log.Printf("%v is building...\n", b.specs["pw"])
			time.Sleep(time.Second * 3)
		}
	}()
	err = cmd.Wait()
	b.waiting = false
}
