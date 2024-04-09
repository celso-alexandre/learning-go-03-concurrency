package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMsg(m string) {
	defer wg.Done() // Debits the wg "Add" counter by 1 after the function completes (even if it panics)
	msg = m
}

func main() {
	msg = "Hello, World!"

	wg.Add(2)
	go updateMsg("Hello, Go!")
	go updateMsg("Hello, Concurrency!")
	wg.Wait() // Blocks until the wg "Add" counter is 0

	fmt.Println(msg)
}
