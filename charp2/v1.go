package main

import "fmt"

var x int
var b byte
var boo bool

var f float32

var com complex64
var xa = 200

func main() {
	//print(111)

	var y,z int
	var a,s string

	fmt.Println(x,b,boo,f,com)

	fmt.Println(y,z,a,s)

	println(&xa,xa)

	var xa = 200
	print(&xa,xa)

	print("=================")

	ac := 100;
	print(&ac,ac)
	ac = 200;
	print(&ac,ac)
}
