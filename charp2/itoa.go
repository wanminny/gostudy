package main

const (

	x = iota
	y
	z
	o = iota
	p
	q


)

const yy = 0x20000
func main() {

	print(x,y,z)
	print(o,p,q)
	println(yy)

}
