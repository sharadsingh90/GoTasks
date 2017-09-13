package main

import (
	"fmt"
	"sync"
)

func writer(reduceQ chan string, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case p := <-reduceQ:
			fmt.Println(p)
		case <-quit:
			return
		}
	}
}
