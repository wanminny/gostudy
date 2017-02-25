package main

import (
	"fmt"
	"reflect"
	"log"
)

type X int
type Y int


func main() {

	var a,b X =  1,2
	var c Y = 3

	fmt.Print(reflect.TypeOf(a) == reflect.TypeOf(b))

	fmt.Print(reflect.Kind(a) == reflect.Kind(b))
	fmt.Print(reflect.Kind(b) == reflect.Kind(c))

	print("=========================")
	//fmt.Print(reflect.Kind(a))
	fmt.Print(reflect.TypeOf(a).Name())
	fmt.Print(reflect.TypeOf(a).Kind())

	log.Print("aaaaaaaaaaaaa")


}
