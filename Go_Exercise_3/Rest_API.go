package main

import (
	"fmt"
	"net/http"
)

var GLOBAL_MAP map[string]string

func initialize() {
	GLOBAL_MAP = make(map[string]string)
}

/*
REQUESTS
//dump request "http://localhost:9090/dump?key=id&value=name"
//fetch request "http://localhost:9090/fetch?key=id"
//count request "http://localhost:9090/count"
//delete request "http://localhost:9090/delete?key=id"
*/
func handlerdelete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	fmt.Println("key is", key)
	delete(GLOBAL_MAP, key)
	out_string := "key: " + key + " deleted."
	fmt.Fprintln(w, out_string)

}
func handlercount(w http.ResponseWriter, r *http.Request) {
	size := len(GLOBAL_MAP)
	fmt.Fprintln(w, size)
}
func handlerretrieve(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	fmt.Fprintln(w, "%s", GLOBAL_MAP[key])
}
func handlerdump(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	fmt.Println("Parameters are", key, "===", value)
	GLOBAL_MAP[key] = value
	out_string := "key: " + key + " value: " + value + " saved."
	fmt.Fprintln(w, out_string)
}
func main() {
	initialize()
	http.HandleFunc("/fetch", handlerretrieve)
	http.HandleFunc("/delete", handlerdelete)
	http.HandleFunc("/count", handlercount)
	http.HandleFunc("/dump", handlerdump)
	http.ListenAndServe(":9090", nil)
}
