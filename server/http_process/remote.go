package http_process

import(
	"fmt"
	"os"
	"net/http"
	"io"
	//"./command"
)


func Command_handler( w http.ResponseWriter, r *http.Request ) {

	sh_file, _, err := r.FormFile( "sh_file" )

	if err != nil {
		fmt.Fprintf( w, err.Error() )
		os.Exit( 0 )
	}

	id := r.Header.Get( "ID" )

	file_name := "./order/" + id + ".txt"
	
	file, err := os.OpenFile( file_name, os.O_WRONLY|os.O_CREATE, 0666 )
	
	defer file.Close()

	if err != nil {
		fmt.Fprintf( w, err.Error() )
		os.Exit( 0 )
	}

	_, err = io.Copy( file, sh_file )

	if err != nil {
		fmt.Fprintf( w, err.Error() )
		os.Exit( 0 )
	}

	//fmt.Fprintln( file, sh_com )
	fmt.Fprintf( w, "sh_command upload success!!!" )
}

