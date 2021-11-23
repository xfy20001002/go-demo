package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var id int //记录上一次打印的id
var expireTime time.Time

func generateBarCodeString(Model, Type, Vendor string, buy int, productTime time.Time) string {
	//条形码从001开始
	if id == 0 {
		id = 1
	}
	if expireTime.IsZero() {
		year, month, day := productTime.Date()
		expireTime = time.Date(year, month, day, 23, 59, 59, 59, time.Local)
	} else {
		if productTime.After(expireTime) {
			id = 1
		}
	}
	SProductedTime := productTime.Format("2006-01-02")
	year := SProductedTime[2:4]
	month := SProductedTime[5:7]
	day := SProductedTime[8:10]
	Sid := formatId(id)
	s := fmt.Sprintf("VTH%s%s-%s-%s%s%s%s%s", Model, Type, Vendor, strconv.Itoa(buy), year, month, day, Sid)
	id = id + 1
	return s
}

func formatId(id int) string {
	//将id格式化
	var sid string
	if id < 10 {
		sid = fmt.Sprintf("00%s", strconv.Itoa(id))
	} else {
		if id < 100 {
			sid = fmt.Sprintf("0%s", strconv.Itoa(id))
		} else {
			sid = strconv.Itoa(id)
		}
	}
	return sid
}

func main() {
	//创建文件
	file, error := os.Create("./barCode.txt")
	if error != nil {
		fmt.Println(error)
	}
	for i := 0; i < 100; i++ {
		//写入byte的slice数据
		barCode := generateBarCodeString("011", "010", "1001", 2, time.Now())
		file.Write([]byte(barCode + "\r\n"))
	}
	file.Close()
}
