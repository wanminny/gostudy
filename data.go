package main

import "fmt"

func main() {
	var dynamic = make([]int,0,5)

	fmt.Print(cap(dynamic),len(dynamic))

	for i :=0;i< 8;i++{

		//fmt.Println(i)
		dynamic = append(dynamic,i)

	}
	print("++++++++++++++++++++")
	print(len(dynamic))
	for k,v := range dynamic{
		fmt.Println(k,"===========>",v)
	}
	print(dynamic[0])

}