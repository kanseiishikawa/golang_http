package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)


func main() {
	url := "http://13.58.72.57/apple"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println( "Error:http" )
		fmt.Println( err )
		return
	}
	
	req.Header.Set( "Hello", "App" )

	client := new(http.Client)
	resp, _ := client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

}
