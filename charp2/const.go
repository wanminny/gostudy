package main

import (
	"fmt"
	//"unsafe"
)

const A = 13
const NAME  = "baby"

func main() {
	fmt.Println(A,NAME)


	const AA = 44.5

	fmt.Print(AA)

	//const ptrSize = unsafe.Sizeof(uintptr(0))

	const  (
		ZZZ = "DDDD"
		//pt = unsafe.Sizeof(uintptr(0))
	)
	var (
		x = 1
	)

	const len1 = len("HELLO")

	print(x,len1,ZZZ)
}
