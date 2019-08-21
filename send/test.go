package main

import (
	"net/http"
	"io/ioutil"
	"io"
	"encoding/json"
	"fmt"
	"os"
	"mime/multipart"
	"bytes"
)

type json_data struct {
	IP string `json:"IP"`
}

func main() {

	var url = connect_url()
	url += "/target/search_sh"

	req, err := http.NewRequest( "Get", url, nil )

	if err != nil {
		fmt.Println( "Error:http" )
		fmt.Println( err )
		os.Exit( 0 )
	}

	req.Header.Set( "ID", "1" )
	//req.Header.Set( "user", "kansei" )
	
	client := new( http.Client )
	resp, err := client.Do( req )

	if resp != nil {
		defer resp.Body.Close()
		var byteArray, _ = ioutil.ReadAll( resp.Body )
		fmt.Println( string( byteArray ) )
	} else {
		fmt.Println( "Send Error")
		fmt.Println( err )
	}
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

func send_file_pre( file_name string ) ( bytes.Buffer, string, error ) {
	var buf bytes.Buffer
	w := multipart.NewWriter( &buf )

	file, err := os.Open( file_name )

	if err != nil {
		return buf, "Error", err
	}

	defer file.Close()

	fw, err := w.CreateFormFile( "sh_file", file_name )

	if err != nil {
		return buf, "Error", err
	}

	_, err = io.Copy( fw, file )

	if err != nil {
		return buf, "Error", err
	}

	w.Close()

	return buf, w.FormDataContentType(), err
}
