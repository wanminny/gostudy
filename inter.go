package main

import "fmt"

type Printer interface {
	Print()
}


type User struct {
	name string
	age byte
}

func (u User)Print()  {
	fmt.Println(u)

}

func main() {
	var u1 User
	u1.name = "wan"
	u1.age = 30

	var p Printer
	p = u1
	p.Print()
}
