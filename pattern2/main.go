package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("Main")
	inputQ, mapQ, reduceQ := make(chan string), make(chan string), make(chan string)
	wgRead, wgMapper, wgWriter := new(sync.WaitGroup), new(sync.WaitGroup), new(sync.WaitGroup)
	wgRead.Add(1)
	go reader(inputQ, mapQ, wgRead)
	wgMapper.Add(1)
	go normalize(mapQ, reduceQ, wgMapper)
	wgWriter.Add(1)
	go writer(reduceQ, wgWriter)
	for i := 0; i < 10; i++ {
		inputQ <- strconv.Itoa(i)
	}
	close(inputQ)
	wgRead.Wait()
	close(mapQ)
	wgMapper.Wait()
	close(reduceQ)
	wgWriter.Wait()
}
