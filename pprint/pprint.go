package main

import (
	"fmt"
	"reflect"

	"../logging"
)

var log interface{}

var test string

//keyword set identifier
type keyword struct {
	identifier string
	startWord  string
	EndWord    string
}

//Pprint struct
type Pprint struct {
	formatList []keyword
	indent     int
}

//NewPprint construct
func NewPprint() *Pprint {
	pprint := new(Pprint)
	mapKey := keyword{"map", "{", "}"}
	ArrayKey := keyword{"array", "[", "]"}
	pprint.formatList = append(pprint.formatList, mapKey)
	pprint.formatList = append(pprint.formatList, ArrayKey)
	return pprint
}

//pprint print list slice
func (p *Pprint) pprint(data interface{}) {
	t := reflect.TypeOf(data).Kind()
	if t == reflect.Map {
		p.PrintMaps(data)
	} else if t == reflect.Slice {
		fmt.Println("[slice]")
		p.PrintSlices(data)
	} else if t == reflect.Array {
		fmt.Println("[Array]")
		p.PrintArrays(data)
	} else {
		fmt.Println("[default]")
	}
	Debug(t)
}

//PrintMaps print map
func (p *Pprint) PrintMaps(data interface{}) {

}

//PrintArrays print map
func (p *Pprint) PrintArrays(data interface{}) {

}

//PrintSlices print slice
func (p *Pprint) PrintSlices(data interface{}) {

}

// Debug temp
func Debug(v interface{}) {
	logg, _ := log.(logging.Logging)
	logg.SetFormat("%#v\n")
	// logg.GetLogger()
	logg.Debug(v)
}

func main() {
	log := logging.NewLogging()
	// log.Debug("temp")
	log.Test()
	m := map[string]string{"foo": "0", "hello": "0"}
	// m := map[string]int{"foo": 0, "hello": 0}
	s := []int{}
	var b [3]int
	pprint := NewPprint()
	pprint.pprint(m)
	pprint.pprint(s)
	pprint.pprint(b)

	// n := map[int]string{1: "bar", 3: "world"}
	// n := [][]map[string]string
	// matrix := [][]string{{[]string{"0"}, []string{"0"}}}
	// Debug(reflect.TypeOf(m).Kind() == reflect.Map)
}
