package fire

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
	"sync"
)

func Run(filepath string) {
	var c Config
	c.Parce(filepath)
	var wg sync.WaitGroup
	wg.Add(len(c.Actions) * c.Target.Users)
	for i := 0; i < c.Target.Users; i++ {
		for _, action := range c.Actions {
			go action.Execute(&wg)
		}
	}
	wg.Wait()
}

type Action struct {
	Type    string `yaml:"type,omitempty"`
	Command string `yaml:"command,omitempty"`
}

type Result struct {
	Output []byte
	Error  error
	Code   int
}

func (a *Action) Execute(w *sync.WaitGroup) {
	defer w.Done()
	command := strings.Split(a.Command, " ")
	cmd := exec.Command(command[0], command[1:]...)

	r, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	done := make(chan struct{})

	// Create a scanner which scans r in a line-by-line fashion
	scanner := bufio.NewScanner(r)

	// Use the scanner to scan the output line by line and log it
	// It's running in a goroutine so that it doesn't block
	go func() {
		// Read line by line and process it
		for scanner.Scan() {
			line := scanner.Text()
			log.Println(line)
		}
		// We're all done, unblock the channel
		done <- struct{}{}
	}()

	// Start the command and check for errors
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}
	// Wait for all output to be processed
	<-done

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		log.Println(err)
	}
}
