package main

import "fmt"

func makeInt() []int {
	tmp := make([]int,0,10)

	tmp = append(tmp,100)

	return tmp
}

func makeMap() map[string]int  {
	tmp := make(map[string]int)

	tmp["aaa"] = 11
	return tmp;
}

func main() {
	a := makeInt()

	b := makeMap()

	fmt.Print(a,b)

	var aa = "abcdefg"
	for i,s := range aa{
		fmt.Println(i,string(s))
	}

	var bb = []string{"c","d","e"}
	for i,s := range bb{
		fmt.Println(i,string(s))
	}

}