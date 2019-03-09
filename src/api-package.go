package main

import (
	"io"
	"log"
	"memcached"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/get/", func(w http.ResponseWriter, req *http.Request) {
		parts := strings.Split(req.URL.Path, "/")
		io.WriteString(w, memcached.Get(parts[2]))
	})
	http.HandleFunc("/set/", func(w http.ResponseWriter, req *http.Request) {
		parts := strings.Split(req.URL.Path, "/")
		memcached.Set(parts[2], parts[3])
		io.WriteString(w, "OK")
	})
	http.HandleFunc("/delete/", func(w http.ResponseWriter, req *http.Request) {
		parts := strings.Split(req.URL.Path, "/")
		memcached.Del(parts[2])
		io.WriteString(w, "OK")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
