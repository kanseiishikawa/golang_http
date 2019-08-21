package main

import (
	//"os/exec"
	"io/ioutil"
	"fmt"
)

func main() {
	dir := "./http_process/"
	files, _ := ioutil.ReadDir( dir )

	fmt.Println( len( files ) )

}
