package main

import (
	"./target_api"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

type json_data struct {
	IP string `json:"IP"`
}


func main() {
	url := connect_url()

	if url == "None" {
		os.Exit( 0 )
	}

	first_req := target_api.First_send( url )

	if first_req == "None" {
		os.Exit( 0 )
	}

	if !write_id( first_req ) {
		os.Exit( 0 )
	}

}

func write_id( id string ) bool {
	file_name := "id_name.txt"

	file, err := os.OpenFile( file_name, os.O_WRONLY|os.O_CREATE, 0666 )
	defer file.Close()

	if err != nil {
		return false
	}

	fmt.Fprintln( file, id )
	return true
}

func connect_url() string {
	url := "http://"
	raw, err := ioutil.ReadFile("../config.json")

	if err != nil {
		return "None"
	}

	var data json_data

	json.Unmarshal( raw, &data )

	url += data.IP

	return url
}

