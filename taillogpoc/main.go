package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	upload, sig, quit := make(chan string), make(chan os.Signal), make(chan bool)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	wg1, wg2 := new(sync.WaitGroup), new(sync.WaitGroup)
	wg2.Add(1)
	go Process1(upload, quit, wg2)
	wg1.Add(1)
	go Process2(upload, quit, wg1)

	go func() {
		<-sig
		for i := 2; i > 0; i-- {
			quit <- true
		}
	}()
	wg1.Wait()
	wg2.Wait()
	fmt.Println("Bye ")
}

func Process2(ch1 chan string, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	cfg := tail.Config{Follow: true, ReOpen: true}
	t, err := tail.TailFile("sftp_ftp.log", cfg)
	if err != nil {
		log.Fatalln("TailFile failed - ", err)
	}
	for {
		select {
		case ev := <-t.Lines:
			ch1 <- ev.Text
		case <-quit:
			return
		}
	}
}

func Process1(ch2 chan string, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case line, ok := <-ch2:
			if !ok {
				return
			}
			fmt.Println(line)
		case <-quit:
			return
		case <-time.After(time.Second):
		}
	}

}

