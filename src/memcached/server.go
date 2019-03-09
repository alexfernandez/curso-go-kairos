package memcached

import (
	"io"
	"log"
	"net/http"
	"strings"
)

var server *http.Server

func Start() {

	http.HandleFunc("/get/", func(w http.ResponseWriter, req *http.Request) {
		parts := strings.Split(req.URL.Path, "/")
		io.WriteString(w, Get(parts[2]))
	})
	http.HandleFunc("/set/", func(w http.ResponseWriter, req *http.Request) {
		parts := strings.Split(req.URL.Path, "/")
		Set(parts[2], parts[3])
		io.WriteString(w, "OK")
	})
	http.HandleFunc("/delete/", func(w http.ResponseWriter, req *http.Request) {
		parts := strings.Split(req.URL.Path, "/")
		Del(parts[2])
		io.WriteString(w, "OK")
	})
	server = &http.Server{
		Addr: ":8080",
	}
	log.Fatal(server.ListenAndServe())
}

func Stop() {
	server.Close()
}
