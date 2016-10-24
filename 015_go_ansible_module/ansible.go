//usr/bin/env ORIG=$0 /bin/bash -c 'tmp=$(mktemp).go;cp $ORIG $tmp;go run $tmp $@';exit
package main

/*
git clone git://github.com/ansible/ansible.git --recursive
source ansible/hacking/env-setup

ansible/hacking/test-module -m ./ansible.go
*/
import (
	"encoding/json"
	"fmt"
	"log"
)

type output struct {
	Date string
}

func main() {
	t := output{Date: "test"}
	o, e := json.Marshal(t)
	if e != nil {
		log.Panic(e)
	}

	fmt.Println(string(o))
}
