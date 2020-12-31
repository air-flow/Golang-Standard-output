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
	switch v := data.(type) {
	case map[string]string:
		// fmt.Println("map")
		fmt.Println(reflect.TypeOf(v))
	case []string:
		// fmt.Println("array")
		fmt.Println(reflect.TypeOf(v))
	case int:
		// fmt.Println(v, "int")
		fmt.Println(reflect.TypeOf(v))
	default:
		fmt.Println(reflect.TypeOf(v))
		logg, _ := log.(logging.Logging)
		logg.Debug(v)
		// fmt.Printf("%v[default]\n", v)
		fmt.Println(test)
	}
}

//PrintMaps print map
func (p *Pprint) PrintMaps() {

}

//PrintArrays print map
func (p *Pprint) PrintArrays() {

}

func main() {
	//ToDo map array slice Judgment
	log := logging.NewLogging()
	log.Debug("temp")
	// logging.Debug(1)
	// m := map[string]string{"foo": "bar", "hello": "world"}
	// n := map[int]string{1: "bar", 3: "world"}
	// n := [][]map[string]string
	// matrix := [][]string{{[]string{"0"}, []string{"0"}}}
	// test := map[string][2]string
	// a := 1
	// pprint := NewPprint()
	// pprint.pprint(m)
	// pprint.pprint(n)
	// pprint.pprint(a)
	// if reflect.TypeOf(0) == int {
	// 	fmt.Println(0)
	// } else {
	// 	fmt.Println(1)
	// }
	// mm := make(map[stri	ng]string, 0)
	// m1 := map[string]int{
	// 	"a": 2,
	// 	"b": 2,
	// }
	// m2 := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// }
	// logging.Debug(reflect.DeepEqual(m1, m2))
	// logging.Debug()
	// fmt.Println(reflect.DeepEqual(m1, m2))
	// fmt.Println(reflect.DeepEqual(&m, &mm))
}
