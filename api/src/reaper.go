package main

import (
	"os/exec"
	"strings"
)

// Kills any containers than have been running longer than 12 hours
func killContainers() int {

	output, err := exec.Command("docker", "ps", "-a", "--format", "\"{{.ID}} {{.RunningFor}}\"").Output()
	if err != nil {
		panic(err)
	}

	containers := strings.Split(strings.Trim(string(output), " \"\n"), "\n")
  numDead := 0

	for i := 0; i < len(containers); i++ {
		timeString := strings.Split(strings.Trim(containers[i], "\" "), " ")
		containerID := timeString[0]
		if timeString[1] == "12" && timeString[2] == "hours" {
			cmd := exec.Command("docker", "container", "rm", "-f", containerID)
			err := cmd.Run()
			if err != nil {
				panic(err)
			}
      numDead += 1
		}
	}
  return numDead
}
