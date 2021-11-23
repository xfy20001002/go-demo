package main

import (
	"fmt"
)

func main() {
	var sn string = "V0101012111001"
	//sntype := reflect.TypeOf(sn[1])
	fmt.Printf("%d", sn[1]-'0')
}
