package main

import (
	"errors"
	// "fmt"
	"log"
	"os"
)

// generic command
type Command interface {
	Run() error
	DryRun() string
}

// instance of success
type GoodCommand string

func (cmd GoodCommand) Run() error {
	return nil
}

func (cmd GoodCommand) DryRun() string {
	return string(cmd)
}

// instance of failure
type BadCommand struct {
	Command
}

func (BadCommand) Run() error {
	return errors.New("Bad Time")
}

func TryCmd(cmd Command, logger *log.Logger) {
	logger.Println("Trying:", cmd.DryRun())
	err := cmd.Run()
	if err != nil {
		logger.Println("Error:", err)
	}
}

func main() {
	/* flags:
			1 - date
		  2 - time (s)
	    4 - time (ms)
	    8 - source line (complete path)
	    16 - source line (just file)
	*/
	logger := log.New(os.Stderr, "DEMO ", 4+16)

	good := GoodCommand("ls /")
	bad := BadCommand{GoodCommand("rm -rf /")}

	TryCmd(good, logger)
	TryCmd(bad, logger)
}
