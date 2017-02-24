package main

import (
	"os"
	"fmt"
)

func loadFile(fileName string)  {
	file,err :=os.Open(fileName)
	if err!=nil{
		panic(err)
	}else {
		fmt.Print(file)
	}

	defer file.Close()
}
func main() {
	loadFile("a1rgs.go")
}
