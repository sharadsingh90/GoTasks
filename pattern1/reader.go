package main

import (
	"sync"
)

func reader(inputFiles chan string, reduceQ chan string, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case p := <-inputFiles:
			reduceQ <- p
		case <-quit:
			return
		}
	}
}
