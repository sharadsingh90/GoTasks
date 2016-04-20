package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
REQUEST
http://localhost:9090 -d '%CONTENT%'
*/
func sampleprogram(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	b := md5.Sum(body)
	pass := hex.EncodeToString(b[:])
	fmt.Fprintln(w, "Md5 of the content", string(body), "is= ", pass)
}
func main() {
	http.HandleFunc("/md5sum", sampleprogram)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
