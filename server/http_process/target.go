package http_process

import(
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"strconv"
	//"./command"
)


func First( w http.ResponseWriter, r *http.Request ) {

	OS := r.Header.Get( "OS" )
	user := r.Header.Get( "user" )

	dir := "./data_storage/"
	my_id := next_id( dir )
	file_name := dir + strconv.Itoa( my_id )+ ".txt"
	
	file, err := os.OpenFile( file_name, os.O_WRONLY|os.O_CREATE, 0666 )
	defer file.Close()

	if err != nil {
		fmt.Fprintf( w, err.Error() )
		return
	}

	fmt.Fprintln( file, "OS " + OS )
	fmt.Fprintln( file, "User " + user )
	
	fmt.Fprintf( w, strconv.Itoa( my_id ) )
}

func next_id( dir string ) int {
	files, _ := ioutil.ReadDir( dir )
	return len( files ) + 1
}
