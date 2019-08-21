package request

import (
	"net/http"
	"io/ioutil"
	"io"
	"mime/multipart"
	"bytes"
	"os"
)

func Send_file( url string, file_name string, ID string ) string {

	sh, form_type, err := send_file_pre(  file_name )

	if err != nil {
		return err.Error()
	}

	req, err := http.NewRequest( "Post", url, &sh )

	if err != nil {
		return err.Error()
	}

	req.Header.Set( "Content-Type", form_type )
	req.Header.Set( "ID", ID )
	
	client := new( http.Client )
	resp, err := client.Do( req )

	if resp != nil {
		defer resp.Body.Close()
		var byteArray, _ = ioutil.ReadAll( resp.Body )
		return string( byteArray )
	} else {
		return err.Error()
	}
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
