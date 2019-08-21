package main

import(
	"fmt"
	//"os"
	"net/http"
	//"io"
	"./http_process"
)

func main() {
	http.HandleFunc("/remote/sh_command", http_process.Command_handler )//シェルのファイルを送って保存する
	
	http.HandleFunc("/target/first_connect", http_process.First )

	fmt.Println("start")
	err := http.ListenAndServe(":80", nil)

	if err != nil{
		fmt.Println( err )
	}
}
