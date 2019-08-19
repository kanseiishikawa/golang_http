package main

import (
	"os/exec"
	"fmt"
)

func main() {
	var out, err = exec.Command( "sh", "hello.sh" ).Output()

	fmt.Println( err )
	fmt.Println( string( out ) )
}
