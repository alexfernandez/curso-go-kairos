package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

var mem map[string]string = make(map[string]string)
var mux sync.Mutex

func get(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	key := parts[2]
	log.Println(key)
	mux.Lock()
	io.WriteString(w, mem[key])
	mux.Unlock()
}

func set(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	key := parts[2]
	value := parts[3]
	log.Println(key + "-" + value)
	mux.Lock()
	mem[key] = value
	mux.Unlock()
	io.WriteString(w, "OK")
}

func del(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	key := parts[2]
	delete(mem, key)
	io.WriteString(w, "OK")
}

func main() {

	http.HandleFunc("/get/", func(w http.ResponseWriter, req *http.Request) {
		get(w, req)
	})
	http.HandleFunc("/set/", func(w http.ResponseWriter, req *http.Request) {
		set(w, req)
	})
	http.HandleFunc("/delete/", func(w http.ResponseWriter, req *http.Request) {
		del(w, req)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
