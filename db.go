package main


import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func checkErr(err error)  {
	if err != nil{
		fmt.Fprint(os.Stderr,"Error : %s",err.Error())
		os.Exit(1)
	}
}

//go mysql 中有数据库连接池的默认实现

func main() {
	con,err := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/python")

	if err!=nil {
		print(err)
	} else {
		fmt.Print("ok!")
		rows,err :=con.Query("SELECT  * FROM info")
		checkErr(err)
		fmt.Print(222222222)
		if err!=nil{
			fmt.Print(err)
			//panic(err)
		}else{
			for rows.Next(){
				var uid int
				var url string
				var content string
				var title string
				rows.Columns()
				rows.Scan(&uid,&url,&title,&content)
				fmt.Println(uid,url)
				fmt.Println(title)
				fmt.Println(content)
			}
		}


	}


	defer con.Close()

}
