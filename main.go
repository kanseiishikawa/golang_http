package main

import(
	"fmt"
	"net/http"
)

func hello_handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, World")
	fmt.Fprintf(w, "Hello" )
}

func apple_handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, World")
	var req_header = r.Header.Get( "Hello" )

	if req_header == "Apple" {
		fmt.Fprintf(w, "Apple_handler" )
	} else {
		fmt.Fprintf(w, "No_Apple" )
	}
}

func main() {
	http.HandleFunc("/hello", hello_handler) // ハンドラを登録してウェブページを表示させる
	http.HandleFunc("/apple", apple_handler) // ハンドラを登録してウェブページを表示させる
	fmt.Println("start")
	err := http.ListenAndServe(":80", nil)

	if err != nil{
		fmt.Println( err )
	}
}
