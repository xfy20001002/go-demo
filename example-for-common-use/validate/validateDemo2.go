package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

//required表明不能使用默认值
type User struct {
	Name string `validate:"required"`
	Age  int    `validate:"required"`
}

func main() {
	validate := validator.New()

	u1 := User{Name: "lidajun", Age: 18}
	err := validate.Struct(u1)
	fmt.Println(err)

	u2 := User{Name: "", Age: 0}
	err = validate.Struct(u2)
	fmt.Println(err)
	//result:Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag
	//result:Key: 'User.Age' Error:Field validation for 'Age' failed on the 'required' tag
}
