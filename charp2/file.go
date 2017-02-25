package main

import (
	"fmt"
	"os"
	"io"
	"strings"
)

func main() {

	name := "wanmin"

	str := fmt.Sprint(
		` <html><head><title>go file</title></head><body>
		asdfsdfsadfa
		<h1> `+ name + `
		</h1>
		</html>
		`)

	//print(str)
	h,err := os.Create("index.html")
	if err != nil{
		fmt.Print(err)
	}

	io.Copy(h,strings.NewReader(str))

	defer h.Close()

}
