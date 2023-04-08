package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func HomeDir() string {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("home directory not found: ", err)
	}
	return home
}

func ValidateCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)

	if err != nil {
		log.Fatal("there waa an error: ", err)
		return false
	}
	return true
}

func CreateDirectory(name string, dir string) string {
	homeDir := HomeDir()
	directory := homeDir

	// use user specified directory
	if len(dir) != 0 {
		directory = fmt.Sprintf("%s/%s", homeDir, dir)
	}

	cmd := exec.Command("mkdir", name)
	cmd.Dir = directory
	err := cmd.Run()

	if err != nil {
		log.Fatal("there was an error creating directory: ", err)
	}

	return fmt.Sprintf("%s/%s", directory, name)
}

func ExecTime() func() {
	start := time.Now()
	return func() {
		fmt.Printf("=> Time spent: %v\n", time.Since(start))
	}
}
