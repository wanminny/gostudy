package main

import (
	"net/http"
	"io"
)

func main() {

	//http.HandlerFunc("index")
	http.HandleFunc("/",index)
	//http.Handle("/",http.HandlerFunc(index))
	http.ListenAndServe(":8000",nil)
}

func index(w http.ResponseWriter, r *http.Request)   {

	io.WriteString(w,"hello bitch !")
}