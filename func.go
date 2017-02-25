package main

import (
	"errors"
	//"os"
	"fmt"
)

func returnFuc(x int) func()  {

	return func() {
		print(x)
	}
}

func div(a,b int) (int,error) {
	if b ==0 {
		return 0,errors.New("diveed by zero!!")
	}else{
		return a/b,nil
	}

}

func main() {

	data,err := div(2,0)
	if err !=nil{
		fmt.Println(err)
		//print(err)
		//os.Exit(1)
	}
	print(data)

	print("================")
	f := returnFuc(100)
	f()

}
