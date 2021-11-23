package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 我们一般使用系统时间的不确定性来进行初始化
	rand.Seed(time.Now().Unix())
	//生成10个小于10的整数
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(10), " ")
	}
}
