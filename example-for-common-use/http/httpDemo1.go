package main

import (
	"net/http"
)

type Server struct {
}

func handle1(w http.ResponseWriter, r *http.Request) {
	s := "root"
	w.Write([]byte(s))
}

func handle1_a(w http.ResponseWriter, r *http.Request) {
	s := "root/a"
	w.Write([]byte(s))
}
func handle1_b(w http.ResponseWriter, r *http.Request) {
	s := "root/b"
	w.Write([]byte(s))
}

func main() {
	http.HandleFunc("/", handle1)
	http.HandleFunc("/a", handle1_a)
	http.HandleFunc("/b", handle1_b)
	http.ListenAndServe("localhost:8888", nil)
}
