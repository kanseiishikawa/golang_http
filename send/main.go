package main

import (
	"./request"
	"bufio"
    "fmt"
    "os"
	"io/ioutil"
	"encoding/json"
)

type json_data struct {
	IP string `json:"IP"`
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println( "何を行いますか?" )
	fmt.Println( "1 shスクリプトの送信" )

	re := nextLine()

	if re == "1" {
		send_sh()
	} else {
		fmt.Println( "そのような指示はありません" )
	}
}

func send_sh() {
	url := connect_url()
	url += "/remote/sh_command"

	fmt.Println( "ファイル名を入力してください" )
	file_name := nextLine()

	fmt.Println( "IDを入力してください" )
	ID := nextLine()

	req := request.Send_file( url, file_name, ID )
	fmt.Println( req )
}

func connect_url() string {
	var url = "http://"
	raw, err := ioutil.ReadFile("../config.json")

	if err != nil {
		fmt.Println( "json_file read Error" )
		fmt.Println( err )
		os.Exit( 1 )
	}

	var data json_data

	json.Unmarshal( raw, &data )

	url += data.IP

	return url
}

func nextLine() string {
    sc.Scan()
    return sc.Text()
}

