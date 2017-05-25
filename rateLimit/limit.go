package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var a string
var c = make(chan int)

func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	wg := new(sync.WaitGroup)
	send, quit := make(chan int), make(chan bool)
	wg.Add(2)
	go A(send, quit, wg)
	go B(send, quit, wg)
	wg.Wait()
	go func() { // Catch the monitored signals and ask goroutines nicely to quit.
		<-sig
		for i := 2; i > 0; i-- {
			quit <- true
		}
	}()
}

func A(ch chan int, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	r := time.Tick(10 * time.Microsecond)
	for {
		select {
		case <-r:
			ch <- rand.Int()
		case <-quit:
			return
		}
	}
}
func B(ch chan int, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			fmt.Println("Value from channel is", <-ch)
		case <-quit:
			return
		}
	}

}
