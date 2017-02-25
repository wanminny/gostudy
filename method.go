package main


type X int

func (self *X)inc()  {
	*self++
}

func main() {
	var x X
	x.inc()

	print(x)
}
