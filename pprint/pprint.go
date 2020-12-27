package main

import "fmt"

//pprint
type pprint struct {
	format_list map[string][]string
	indent      int
}

//NewPrint construct
func NewPrint() *pprint {
	pprint := new(pprint)
	pprint.format_list["map"] = {"{", "}"}
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
	pprint := new(pprint)
	pprint.pprint(m)
}
