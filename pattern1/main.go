package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	fmt.Println("Main")
	inputQ, mapQ, reduceQ, quit := make(chan string), make(chan string), make(chan string), make(chan bool)
	wgRead, wgMapper, wgWriter := new(sync.WaitGroup), new(sync.WaitGroup), new(sync.WaitGroup)
	wgRead.Add(1)
	go reader(inputQ, mapQ, quit, wgRead)
	wgMapper.Add(1)
	go normalize(mapQ, reduceQ, quit, wgMapper)
	wgWriter.Add(1)
	go writer(reduceQ, quit, wgWriter)
	for i := 0; i < 10; i++ {
		inputQ <- strconv.Itoa(i)
	}
	<-time.After(10 * time.Second)
	go func() {
		for i := 0; i < 3; i++ {
			quit <- true
		}
	}()
	wgRead.Wait()
	wgMapper.Wait()
	wgWriter.Wait()

}
