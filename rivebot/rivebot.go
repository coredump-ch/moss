package rivebot

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

// Structs

type Rivebot struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

type Query struct {
	Message string `json:"message"`
}

type Response struct {
	Status string
	Reply  string
}

// Constructor

func assertFileExists(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalf("No such file or directory: %s", filename)
	}
}

func NewRivebot() *Rivebot {
	var r Rivebot
	var err error
	var filename = "rivebot/rivebot.py"

	assertFileExists(filename)
	r.cmd = exec.Command("python", filename)
	r.stdin, err = r.cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	r.stdout, err = r.cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	return &r
}

// Methods

func (r *Rivebot) Start() {
	log.Print("Starting rivebot...")

	// Start process
	if err := r.cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Create goroutine that waits for the process to finish
	done := make(chan error, 1)
	go func() {
		done <- r.cmd.Wait()
	}()

	// Check whether the process has been started
	select {
	case err := <-done:
		if err == nil {
			log.Fatal("Rivebot process exited without an error.")
		} else {
			log.Fatalf("Rivebot process exited with error %v.", err)
		}
	case <-time.After(time.Second * 2):
		log.Print("Started rivebot...")
	}
}

func (r *Rivebot) Ask(message string) (string, error) {

	// Send query
	var query = Query{message}
	log.Print("Sending query...")
	if err := json.NewEncoder(r.stdin).Encode(query); err != nil {
		log.Fatalf("Could not encode query: %s", err)
	}
	log.Print("Sending __END__...")
	if _, err := r.stdin.Write([]byte("\n__END__\n")); err != nil {
		log.Fatalf("Could not write __END__: %s", err)
	}

	// Read response
	var response Response
	log.Print("Decoding response...")
	if err := json.NewDecoder(r.stdout).Decode(&response); err != nil {
		log.Fatalf("Could not decode response: %s", err)
	}

	// Validate response
	log.Print("Validating response...")
	if response.Status == "" {
		log.Print("Error: Empty status.")
		return response.Reply, errors.New(response.Reply)
	} else if response.Status != "ok" {
		log.Printf("Error: %s", response.Reply)
		return response.Reply, errors.New(response.Reply)
	}

	log.Printf("New answer: %s", response)
	return response.Reply, nil
}
