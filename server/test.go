package main

import (
	//"os/exec"
	"io/ioutil"
	"fmt"
)

func main() {
	dir := "./http_process"
	files, _ := ioutil.ReadDir( dir )

	for _, f := range files {
		if ! f.IsDir() {
			fmt.Println( f.Name() )
		}
	}

}
