package main

import (
	"fmt"
	"os/exec"
)


func main() {
	gitProgram, err := exec.LookPath("git")

	if (err != nil) {
		fmt.Printf("Error trying to find git: %s\n", err)
	} else {
		fmt.Printf("Git found in path: %s\n", gitProgram)
	}
}

