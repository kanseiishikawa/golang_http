package target_api

import (
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
)

func Search_sh( url string, id string ) bool {
	url += "/target/search_sh"

	req, err := http.NewRequest( "Get", url, nil )

	if err != nil {
		return false
	}

	req.Header.Set( "ID", id )

	client := new( http.Client )
	resp, err := client.Do( req )

	if resp != nil {
		defer resp.Body.Close()
		var byteArray, _ = ioutil.ReadAll( resp.Body )

		if string( byteArray ) != "None" {
			return make_sh( string( byteArray ) )
		} else {
			return false
		}
	} else {
		return false
	}
}

func make_sh( sh_data string ) bool {
	file, err := os.OpenFile( "command.sh", os.O_WRONLY|os.O_CREATE , 0666 )

	if err != nil {
		return false
	}

	defer file.Close()
	fmt.Fprintln( file, sh_data )

	return true
}
