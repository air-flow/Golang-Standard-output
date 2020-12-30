package main

import (
	"fmt"
	"log"
	"os"

	"./logging"
)

func main() {
	in := [...]float64{10, 6, 3}
	// var mapEx = map[string][]string{"firstName": ["John"]}
	// temp["temp"]="temp"
	// log.Printf("%#v", mapEx)
	// l := in[0]
	// a := in[1]
	// b := in[2]
	p, _ := os.Getwd()
	os.Chdir(p)
	log.Printf("%#v", in)
	fmt.Printf("%#v", in)
	logging.Test()
	// file_import.TestLog()

}
