package main

import "fmt"

func demo() {
	a := 1
	print(a)
}

func swi() {

	x := 100
	switch {
	case x > 0:
		print(x)
	case x < 0:
		print("minus")
	default:
		println("default!!")

	}

}

func forloop()  {
	for i :=0;i<10;i++{
		println(i)
	}

	for i :=4;i>0;i--{
		println(i)
	}
}
func exp() {
	a := 10
	if a >= 10 {
		println("dddddd", a)
	} else if a > 5 {
		println(a)
	} else {
		println("default!!")
	}
}

func main() {
	var x int
	var s = "string variable"
	var str string

	print(x, str, s)
	fmt.Print(str)
	b := 2
	println(b)
	demo()

	exp()

	swi()
	print("=================\n")
	forloop()

	y :=4
	for y >1{
		print(y)
		y--
	}

	//for{
	//	print("abc")
	//}

	for{
		x :=5
		print(x)
		x--
		if x==0 {
			break
		}
	}
}
