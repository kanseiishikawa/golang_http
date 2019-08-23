package target_api

import (
	"net/http"
	"io/ioutil"
)

func First_send( url string ) string {
	url += "/target/first_connect"

	req, err := http.NewRequest( "Get", url, nil )

	if err != nil {
		return "None"
	}

	t_os, t_name := data_collect()

	req.Header.Set( "OS", t_os )
	req.Header.Set( "user", t_name )

	client := new( http.Client )
	resp, err := client.Do( req )

	if resp != nil {
		defer resp.Body.Close()
		var byteArray, _ = ioutil.ReadAll( resp.Body )
		
		return string( byteArray )
	} else {
		return "None"
	}
}

//今はまだテスト段階なので固定
//最終はshellを使ってdataの収集をしていく
func data_collect() ( string, string ) {
	return "Mac", "test_user"
}
