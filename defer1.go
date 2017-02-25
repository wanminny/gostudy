package main


func deferTest(a,b int){
	defer print("finally......")

	print(a/b)

}

func main() {
	deferTest(1,0)

	//deferTest(1,1)
}

