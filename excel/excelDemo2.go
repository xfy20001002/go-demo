package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx"
)

type Tutorial struct {
	Id      int
	Title   string
	Summary string
	Author  string
}

func main() {
	buf := bytes.NewBuffer(nil)
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")

	tutorials := []Tutorial{
		Tutorial{Id: 1, Title: "Go 入门编程", Summary: "Go 基本语法和使用示例", Author: "学院君"},
		Tutorial{Id: 2, Title: "Go Web 编程", Summary: "Go Web 编程入门指南", Author: "学院君"},
		Tutorial{Id: 3, Title: "Go 并发编程", Summary: "通过并发编程提升性能", Author: "学院君"},
		Tutorial{Id: 4, Title: "Go 微服务开发", Summary: "基于 go-micro 框架开发微服务", Author: "学院君"},
	}

	line := []string{
		"序号",
		"登记日期",
		"项目名称",
		"项目挂网链接",
		"报名截止日期",
	}
	row := sheet.AddRow()
	for _, field := range line {
		cell := row.AddCell()
		cell.Value = field
	}
	for id, company := range tutorials {
		line = []string{
			strconv.Itoa(id + 1),
			strconv.Itoa(company.Id),
			company.Title,
			company.Summary,
			company.Author,
		}
		// 将切片类型行数据写入 csv 文件
		row := sheet.AddRow()
		for _, field := range line {
			cell := row.AddCell()
			cell.Value = field
		}
		file.Write(buf)
	}
	fmt.Println(buf.Bytes())
}
