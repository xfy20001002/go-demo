package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
)

type Tutorial struct {
	Id      int
	Title   string
	Summary string
	Author  string
}

func main() {
	// 创建一个 tutorials.csv 文件
	buf := bytes.NewBuffer(nil)

	// 初始化字典数据
	tutorials := []Tutorial{
		Tutorial{Id: 1, Title: "Go 入门编程", Summary: "Go 基本语法和使用示例", Author: "学院君"},
		Tutorial{Id: 2, Title: "Go Web 编程", Summary: "Go Web 编程入门指南", Author: "学院君"},
		Tutorial{Id: 3, Title: "Go 并发编程", Summary: "通过并发编程提升性能", Author: "学院君"},
		Tutorial{Id: 4, Title: "Go 微服务开发", Summary: "基于 go-micro 框架开发微服务", Author: "学院君"},
	}

	// 初始化一个 csv writer，并通过这个 writer 写入数据到 bytes.buffer缓冲区
	writer := csv.NewWriter(buf)
	for _, tutorial := range tutorials {
		line := []string{
			strconv.Itoa(tutorial.Id), // 将 int 类型数据转化为字符串
			tutorial.Title,
			tutorial.Summary,
			tutorial.Author,
		}
		// 将切片类型行数据写入 csv 文件
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	// 将 writer 缓冲中的数据都推送到 csv 文件，至此就完成了数据写入到 csv 文件
	writer.Flush()

	fmt.Println(buf.Bytes())

}
