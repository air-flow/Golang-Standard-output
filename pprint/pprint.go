package main

import "fmt"

//pprint
type pprint struct {
	format_list []keyword
	indent      int
}

type keyword struct {
	identifier string
	startWord  string
	EndWord    string
}

//NewPrint construct
func NewPrint() *pprint {
	pprint := new(pprint)
	mapKey := keyword{"map", "{", "}"}
	ArrayKey := keyword{"map", "{", "}"}
	pprint.format_list = append(pprint.format_list, mapKey)
	pprint.format_list = append(pprint.format_list, ArrayKey)
	return pprint
}

//pprint print list slice
func (p *pprint) pprint(data interface{}) {
	switch v := data.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("%v\n", v)
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
	pprint := NewPrint()
	pprint.pprint(m)
}
