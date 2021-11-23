package main

import (
	"fmt"
	"reflect"
)

type s struct {
	A string
	B int
}

func main() {
	var a int
	var b uint
	var i interface{}
	c := make([]s, 0)
	d := new(s)
	var strs []string
	var str string
	fmt.Println(reflect.ValueOf(s{}).Kind())  //struct
	fmt.Println(reflect.ValueOf(a).Kind())    //int
	fmt.Println(reflect.ValueOf(b).Kind())    //uint
	fmt.Println(reflect.ValueOf(c).Kind())    //slice
	fmt.Println(reflect.ValueOf(d).Kind())    //ptr
	fmt.Println(reflect.ValueOf(i).Kind())    //invalid
	fmt.Println(reflect.ValueOf(strs).Kind()) //slice
	fmt.Println(reflect.ValueOf(str).Kind())  //string
	//修改结构体内容
	sss := s{A: "hello", B: 12}
	real_s := reflect.ValueOf(&sss).Elem()

	fmt.Println(real_s.Interface())

	real_s.Field(0).SetString("world")
	real_s.Field(1).SetInt(1)
	fmt.Println(real_s.Interface())
	//修改反射结构体数组的内容
	ssss := []s{
		{A: "hello", B: 12},
		{A: "world", B: 11},
		{A: "nihao", B: 13},
		{A: "shijie", B: 15},
	}
	fmt.Println(reflect.ValueOf(&ssss).Kind())        //ptr
	fmt.Println(reflect.ValueOf(&ssss).Elem().Kind()) //slice
	ssssArrayPtr := reflect.ValueOf(&ssss)
	slicev := reflect.ValueOf(&ssss).Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	fmt.Println(slicev.Cap())  //4
	fmt.Println(slicev.Kind()) //slice
	fmt.Println(slicev.Type()) //[]main.s
	//添加新元素
	elemt := slicev.Type().Elem()
	fmt.Println(elemt.Kind()) //struct
	elemp := reflect.New(elemt)
	fmt.Println(elemp.Type())                        //*main.s Kind:ptr
	elemp.Elem().Set(reflect.ValueOf(s{"hhhh", 16})) //设置elemp的值
	fmt.Println(elemp.Interface())                   //&{hhhh 16}
	//将新元素添加到数组
	slicev = reflect.Append(slicev, elemp.Elem())
	ssssArrayPtr.Elem().Set(slicev.Slice(0, slicev.Len()))
	//ssssArrayPtr.Elem().Field(0).SetString("hulahula")
	fmt.Println(ssss) //[{hello 12} {world 11} {nihao 13} {shijie 15} {hhhh 16}]
}
