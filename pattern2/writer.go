package main

import (
	"fmt"
	"sync"
)

func writer(reduceQ chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range reduceQ {
		fmt.Println(p)
	}
}
