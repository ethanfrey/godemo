package main

import (
	//  "errors"
	"log"
	"os"
	"runtime"
	"time"
)

const ms time.Duration = time.Nanosecond * 1000 * 1000

var logger = log.New(os.Stderr, "", 4+16)

func wait(delayMs int64) {
	delay := ms * time.Duration(delayMs)
	time.Sleep(delay)
}

func demo(val int) {
	wait(100)
	logger.Println(val)
}

func cdemo(in chan int, out chan int) {
	logger.Println("goroutine started")
	for {
		data := <-in
		logger.Println(data)
		wait(100)
		out <- data * data
	}
}

func straightPrint() {
	for i := range [5]int{} {
		// this is buggy using closure
		go func() { log.Println(i) }()
		// this works, passes arg to goroutine
		go demo(i)
	}
	wait(300)
}

func chanPrint() {
	count := 5
	in, out := make(chan int, count), make(chan int, count)
	counter := make([]int, 5, 5)

	// start as many works as needed
	go cdemo(in, out)
	go cdemo(in, out)

	// send data
	for i := range counter {
		in <- i
	}
	// read data
	for _ = range counter {
		logger.Println(<-out)
	}
}

func main() {
	// see difference here: GOMAXPROCS=2 go run concurrent2.go
	var numCPU = runtime.GOMAXPROCS(0)
	var numCPU2 = runtime.NumCPU()
	logger.Println("cores:", numCPU, "/", numCPU2)

	// straightPrint()
	chanPrint()
}
