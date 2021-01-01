package main

import (
	"fmt"

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
	switch v := data.(type) {
	case map[string]string:
		fmt.Println("[map]")
		// fmt.Println(reflect.TypeOf(v))
		Debug(v)
	case map[interface{}]interface{}:
		fmt.Println("[map interface]")
		// fmt.Println(reflect.TypeOf(v))
		Debug(v)
	case map[string]int:
		fmt.Println("[map interface]")
		// fmt.Println(reflect.TypeOf(v))
		Debug(v)
	case []string:
		fmt.Println("[array]")
		// fmt.Println(reflect.TypeOf(v))
		Debug(v)
	case int:
		fmt.Println("[int]")
		// fmt.Println(reflect.TypeOf(v))
		Debug(v)
	default:
		fmt.Println("[default]")
		// fmt.Printf("%v[default]\n", v)
		Debug(v)
		// fmt.Println(test)
	}
}

//PrintMaps print map
func (p *Pprint) PrintMaps() {

}

//PrintArrays print map
func (p *Pprint) PrintArrays() {

}

// Debug temp
func Debug(v interface{}) {
	logg, _ := log.(logging.Logging)
	logg.SetFormat("%#v\n")
	// logg.GetLogger()
	logg.Debug(v)
}

func main() {
	//ToDo map array slice Judgment
	log := logging.NewLogging()
	// log.Debug("temp")
	log.Test()

	m := map[string]int{"foo": 0, "hello": 0}
	// m := map[string]int{"foo": 0, "hello": 0}
	pprint := NewPprint()
	pprint.pprint(m)
	// n := map[int]string{1: "bar", 3: "world"}
	// n := [][]map[string]string
	// matrix := [][]string{{[]string{"0"}, []string{"0"}}}
}
