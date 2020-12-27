package main

import (
	"fmt"
	"reflect"
)

//keyword set identifier
type keyword struct {
	identifier string
	startWord  string
	EndWord    string
}

//pprint
type pprint struct {
	formatList []keyword
	indent     int
}

//NewPprint construct
func NewPprint() *pprint {
	pprint := new(pprint)
	mapKey := keyword{"map", "{", "}"}
	ArrayKey := keyword{"map", "{", "}"}
	pprint.formatList = append(pprint.formatList, mapKey)
	pprint.formatList = append(pprint.formatList, ArrayKey)
	return pprint
}

//pprint print list slice
func (p *pprint) pprint(data interface{}) {
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
		// fmt.Printf("%v[default]\n", v)
		// fmt.Println(v)
	}
}

//PrintMap print map
func (p *pprint) PrintMaps() {

}

//PrintArrays print map
func (p *pprint) PrintArrays() {

}

func main() {
	m := map[string]string{"foo": "bar", "hello": "world"}
	n := map[int]string{1: "bar", 3: "world"}
	a := 1
	pprint := NewPprint()
	pprint.pprint(m)
	pprint.pprint(n)
	pprint.pprint(a)
}
