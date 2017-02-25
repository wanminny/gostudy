package main

import "fmt"

type N int

func (N) test() {
	println("hi")
}

func main() {
	var a N

	a.test()


	name := "Todd McLeod"

	tpl := `
	asdffadsfadsfasdfassdf
	asdfsd
	sdag
	asdg
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)

}