package main

import (
	"fmt"
	"log"
	"reflect"
)

// maptest
// func maptest(data interface{}) {
// m := make(map[interface{}]func())
// m[reflect.TypeOf(data).Kind()] := arrayPrint
// m["array"]()
// }

// arrayPrint test
func arrayPrint(test string) {
	fmt.Println("data")
}

func mapreflect() {
	//ToDo map key reflect value struct
	u := make(map[uint]int)
	data := "ste"
	m := reflect.TypeOf(data).Kind()
	u[0] = 0
	// m[reflect.Array] = 0
	fmt.Println(u)
	// fmt.Println(reflect.TypeOf(data).)
	fmt.Println(reflect.TypeOf(m).Kind())
	fmt.Println(reflect.TypeOf(reflect.Array).Kind())
}
func p(prefix string, val interface{}) {
	rt := reflect.TypeOf(val)
	log.Printf(prefix+"\nvalue: %+v, type: %+v, kind: %+v", val, rt, rt.Kind())
}

func FuncVarTest(in interface{}) {
	fmt.Println(in)
}

type TypeDetails struct {
	identifier  string
	Word        [2]string
	funcName    func(interface{})
	reflectType interface{}
}

func (t *TypeDetails) StructTest(in interface{}) {
	fmt.Println(in)
}

func main() {
	// in := [...]float64{10, 6, 3}
	// var mapEx = map[string][]string{"firstName": ["John"]}
	// temp["temp"]="temp"
	// log.Printf("%#v", mapEx)
	// l := in[0]
	// a := in[1]
	// b := in[2]
	// p, _ := os.Getwd()
	// os.Chdir(p)
	// log.Printf("%#v", in)
	// fmt.Printf("%#v", in)
	// logging.Test()
	// file_import.TestLog()
	// b := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// maptest(b)
	// f := make(map[interface{}]func())
	// f["array"] = arrayPrint
	// f["array"]()
	// mapreflect()
	gv := [1]int{1}
	var x, y interface{}
	rv := reflect.TypeOf(gv).Kind()
	x = reflect.Map
	y = rv.String()
	gv2 := rv.String()
	fmt.Println(gv2)
	fmt.Println(x)
	fmt.Println(y)
	f := StructTest
	f("test")
	// p("Int.", gv2)
}
