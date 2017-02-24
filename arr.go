package main

import "fmt"

func test(arg ...int) {
	for _,v := range arg{
		fmt.Print(v)
	}
}

func main() {
	var arr = [10]int{1,2,3,4,5,6,7,8,9,10}


	arr1 := arr[:5]

	for i:= 0;i<len(arr);i++{
		//fmt.Print(arr[i])
	}

	for _,v:= range arr1{
		fmt.Println(v)
	}

	test(1,3,99)
}
