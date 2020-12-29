package main

import (
	"fmt"
	"log"
	"logging"
)

func main() {
	in := [...]float64{10, 6, 3}
	var mapEx = map[string][]string{"firstName": ["John"]}
	// temp["temp"]="temp"
	log.Printf("%#v", mapEx)
	// l := in[0]
	// a := in[1]
	// b := in[2]
	log.Printf("%#v", in)
	fmt.Printf("%#v", in)
	logging.checkMain()
}
