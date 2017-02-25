package main

import (
	"os"
	"fmt"
)

func main() {
	//data,err := os.Open("/dev/random")
	data,err := os.Open("/demo/go/charp2/http.go")



	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}


	defer data.Close()

	buff := make([]byte,1024)
	stream,_ :=data.Read(buff)

	fmt.Print(stream)
	fmt.Print(buff)
}
