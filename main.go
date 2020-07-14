package main

import (
	"fmt"
	demo "github.com/luedigernet/moduleDemo"
	demo2 "github.com/luedigernet/moduleDemo/v2"
	demo3 "github.com/luedigernet/modules-demo/pkg/demo"
)



func main() {
	result:= demo2.SayHello("Reinhard")
	fmt.Println(result)
	demo.SayHello()
	demo3.Add(1,2)
}
