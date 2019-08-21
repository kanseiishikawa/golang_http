package main

func main() {
	url = connect_url()

	
}

func connect_url() string {
	url := "http://"
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

