package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		获取现在时间
	*/
	//获取现在时间
	time1 := time.Now()
	fmt.Println(time1) //result:2021-11-12 21:17:47.6584803 +0800 CST m=+0.002877401
	//得到当前时间年月日
	//得到当前时间时分秒
	year, month, day := time1.Date()
	hour, min, sec := time1.Clock()
	fmt.Println(year, month, day, hour, min, sec) //result:2021 November 12 21 58 13
	/*
		时间格式转换
	*/
	//字符串转时间
	Stime1 := "2021-03-19 09:23:29"
	time2, _ := time.Parse("2006-01-02 15:04:05", Stime1)
	fmt.Println(time2) //result:2021-03-19 09:23:29 +0000 UTC

	//时间转字符串
	Stime2 := time2.Format("2006-01-02 15:04:05")
	fmt.Println(Stime2) //result:2021-03-19 09:23:29

	/*
		比较时间大小
	*/
	//比较t1是否小于t2
	res1 := time1.Before(time2)
	fmt.Println(res1)

	//比较t1是否大于t2
	res2 := time1.After(time2)
	fmt.Println(res2)

	//比较t1是否与t2相等
	res3 := time1.Equal(time2)
	fmt.Println(res3)

	/*
		增减时间
	*/
	// 1分钟前
	m, _ := time.ParseDuration("-1m")
	m1 := time2.Add(m)
	fmt.Println(m1) // result 2021-03-19 09:22:29 +0000 UTC
	// 8个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := time2.Add(8 * h)
	fmt.Println(h1) // result 2021-03-19 01:23:29 +0000 UTC
	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := time2.Add(d)
	fmt.Println(d1) // result 2021-03-18 09:23:29 +0000 UTC
	// 1分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := time2.Add(mm)
	fmt.Println(mm1) // result 2021-03-19 09:24:29 +0000 UTC
	// 1小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := time2.Add(hh)
	fmt.Println(hh1) // result 2021-03-19 10:23:29 +0000 UTC
	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := time2.Add(dd)
	fmt.Println(dd1) // result 2021-03-20 09:23:29 +0000 UTC

	//当前时间转化为0时0分0秒
	date_utc := time.Date(year, month, day, 0, 0, 0, 0, time1.Location())
	fmt.Println(date_utc)

	//判断时间是否为空
	var t time.Time
	fmt.Println(t)
	fmt.Println(t.IsZero())
}
