package main

import "fmt"

type IFly interface {
	Fly()
}
type Bird struct {

}

func (b *Bird)Fly(){
	fmt.Print("ifly()......")
}

func main() {
	var instance  = new(Bird)
	///下面也可以执行
	//var instance IFly = new(Bird)
	instance.Fly()
}

