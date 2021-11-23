package Demo2

import (
	"net/http"
)

var server http.Server

func handle2(w http.ResponseWriter, r *http.Request) {
	s := "root"
	w.Write([]byte(s))
}
func handle2_a(w http.ResponseWriter, r *http.Request) {
	s := "root/a"
	w.Write([]byte(s))
}
func handle2_b(w http.ResponseWriter, r *http.Request) {
	s := "root/b"
	w.Write([]byte(s))
}
func main() {
	server = http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	http.HandleFunc("/", handle2)
	http.HandleFunc("/a", handle2_a)
	http.HandleFunc("/b", handle2_b)
	server.ListenAndServe()
}
