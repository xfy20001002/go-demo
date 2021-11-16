package test

import (
	"io"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		log.Printf("reponse err ")
	}
}

/*
func main() {
	// 路由与视图函数绑定
	http.HandleFunc("/health-check", HealthCheckHandler)

	// 启动服务,监听地址
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
*/
