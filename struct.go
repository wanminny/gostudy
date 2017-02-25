package main

import "fmt"

//import "github.com/howeyc/fsnotify"

type user struct {
	name string
	age byte
}

type manager struct {
	user
	title string
}

func (m manager)ToString() string{

	return fmt.Sprintf("name:%s , age : %d,title: %s",m.name,m.age,m.title)
}
func main() {

	var m1 manager
	m1.age =30
	m1.name = "sssss"
	m1.title =  "ceo"

	//fmt.Println(m1)
	//没有效果
	fmt.Sprintf("%v",m1)
	//fmt.Print(m1)
	println(m1.ToString())
	//print(m1)


}
