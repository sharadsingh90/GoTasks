package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == 0 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("err is", err)
		}
		str := string(bytes)
		data := []byte(str)
		fmt.Printf("%x", md5.Sum(data))
	} else {
		if len(os.Args) == 1 {
			fmt.Printf("%x", md5.Sum([]byte("")))
		} else {
			arg := os.Args[1]
			dat, err := ioutil.ReadFile(arg)
			if err != nil {
				fmt.Println("err is", err)
			}
			data := []byte(dat)
			fmt.Printf("%x", md5.Sum(data))
		}
	}
}
