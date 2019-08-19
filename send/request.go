package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

type json_data struct {
	IP string `json:"IP"`
}

func main() {

	var url = connect_url()
	url += "/apple"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println( "Error:http" )
		fmt.Println( err )
		return
	}
	
	req.Header.Set( "Hello", "Apple" )

	client := new(http.Client)
	resp, _ := client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}

func connect_url() string {
	var url = "http://"
	raw, err := ioutil.ReadFile("./config.json")

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
