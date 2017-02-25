package main

import (
	"fmt"
	"time"
)

func task(id int)  {

	var a = 's'  //单引号用于golang里单引号只能有一个字符


	for i:=0; i<3; i++{
		fmt.Printf("===%d,===> %d,===%v\n",id,i,a)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go task(1)
	go task(2)

	time.Sleep(time.Second * 5)
}