package test

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"net/http/httptest"
	"testing"
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
func TestMyServer(t *testing.T) {
	server = http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	http.HandleFunc("/", handle2)
	http.HandleFunc("/a", handle2_a)
	http.HandleFunc("/b", handle2_b)
	testserver := httptest.NewServer(server) //错误 server没有实现handler接口 无法模拟服务器进行测试
	expect := httpexpect.New(t, testserver.URL)
	date1 := expect.GET("/").Expect().Status(200).Body()
	println(data1.Row())
	date2 := expect.GET("/a").Expect().Status(200).Body()
	println(data2.Row())
	date3 := expect.GET("/b").Expect().Status(200).Body()
	println(data3.Row())
}
