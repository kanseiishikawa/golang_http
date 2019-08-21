package http_process

import(
	"fmt"
	"net/http"
	"io/ioutil"
)


func Target_sh_send( w http.ResponseWriter, r *http.Request ) {
	id := r.Header.Get( "ID" )
	fmt.Fprintf( w,  search_file( id ) )
}

func search_file( my_id string ) string {
	target_name := my_id + ".txt"
	dir := "./order/"

	files, _ := ioutil.ReadDir( dir )

	for _, f := range files {
		if target_name == f.Name() {
			bytes, _ := ioutil.ReadFile( dir + target_name )
			return string( bytes )
		}
	}

	return "None"
} 


