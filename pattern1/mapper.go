package main

import (
	"sync"
)

func normalize(mapQ chan string, reducerQ chan string, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case p := <-mapQ:
			reducerQ <- p
		case <-quit:
			return
		}
	}
}
