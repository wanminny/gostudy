package main

import (
	"html/template"
	"fmt"
	"os"
)

func init()  {
	print("init hook!!")
}
func main() {


	text,err := template.ParseFiles("/demo/go/charp2/addr.go")

	if err != nil{
		fmt.Print(err)
	}

	err = text.Execute(os.Stdout,nil)

	if err != nil{
		fmt.Println(err)
	}
}
