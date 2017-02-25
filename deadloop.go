package main



func rangFuc(){
	var tmp = []int{1,2,44}
	for k,v := range tmp{
		print(k,v)
	}
}
func main() {

	rangFuc()
	x :=5
	for{
		print(x)
		x--
		if x ==0{
			break
		}
	}
}
