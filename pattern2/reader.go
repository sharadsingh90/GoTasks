package main

import (
	"sync"
)

func reader(inputFiles chan string, reduceQ chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range inputFiles {
		reduceQ <- p
	}
}
