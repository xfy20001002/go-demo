package Demo4

import (
	"net/http"
	"text/template"
)

var server http.Server

func handle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("Demo4.html")
	t.Execute(w, "hello world")
}

func main() {
	server = http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	http.HandleFunc("/", handle)
	server.ListenAndServe()
}
