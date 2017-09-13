package main

import (
	"sync"
)

func normalize(mapQ chan string, reducerQ chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range mapQ {
		reducerQ <- p
	}
}
