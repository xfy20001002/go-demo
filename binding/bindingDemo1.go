package main

import (
	"fmt"
	"net/http"

	"github.com/atomtree/binding"
)

//绑定查询参数
type ProjectBiddingQL struct {
	/*
		Skip  int64    `json:"skip"`
		Limit int64    `json:"limit"`
		Sort  []string `json:"sort"`
		Word  string   `json:"word"`
	*/
	Skip  int64
	Limit int64
	Sort  []string
	Word  string
}

func (o *ProjectBiddingQL) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&o.Skip:  "s",
		&o.Limit: "l",
		&o.Sort:  "o",
		&o.Word:  "k",
	}
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		//url: http://localhost:8888/?s=0&l=10&o=-date&k=hello
		ql := new(ProjectBiddingQL)
		if err := binding.Bind(r, ql); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(ql.Skip)
		fmt.Println(ql.Limit)
		fmt.Println(ql.Sort)
		fmt.Println(ql.Word)
		//res:
		//0
		//10
		//[-date]
		//hello
		rw.Write([]byte("hello world"))
	})
	fmt.Println("test server is running...")

	http.ListenAndServe("localhost:8888", nil)
}
