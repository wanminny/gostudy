package main

import (
	"os"
	"fmt"
	//"time"
	"time"
)

func test()  {
	print(111111)
}
func main(){
	var s string
	for i:=1;i<len(os.Args);i++  {
		//tmp = ''
		s+= os.Args[i]
	}

	fmt.Print(s)
	//time.Sleep(10000000)
	//test()
	go test()
	time.Sleep(1*time.Second)

}
