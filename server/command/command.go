package command

import (
	"os/exec"
)

func Result( com string ) ( string, error ) {
	var err error
	var out = []byte{}
	out, err = exec.Command( "sh", com ).Output()

	return string( out ), err
}
