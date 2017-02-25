package main

func main() {
	var m1 = make(map[string]int)

	m1["baidu.com"] = 1

	print(m1["baidu.com"])

	x,ok := m1["b"]

	print(x,ok)

	delete(m1,"baidu.com")


	for k,v := range m1{
		print(k,v)
	}

	//print(m1["baidu"])
}
