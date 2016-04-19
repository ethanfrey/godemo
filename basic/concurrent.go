package main

import (
	//	"errors"
	"log"
	"os"
	"time"
)

const ms time.Duration = time.Nanosecond * 1000 * 1000

var logger = log.New(os.Stderr, "", 4+16)

func wait(delayMs int64) {
	delay := ms * time.Duration(delayMs)
	time.Sleep(delay)
}

func output(delayMs int64, msg string, out chan int) {
	logger.Println(msg)
	wait(delayMs)
	logger.Println(msg)
	out <- 1
}

func main() {
	c := make(chan int)
	go output(1000, "first", c)
	go output(500, "second", c)
	// must wait for both goroutines to end, or nothing printed (also time.Sleep works)
	<-c
	<-c
}
