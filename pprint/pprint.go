package main

import (
	"fmt"
	"reflect"
	"strings"

	"../logging"
)

var log interface{}

//keyword set identifier
type keyword struct {
	identifier string
	Word       [2]string
}

//Pprint struct
type Pprint struct {
	formatList []keyword
	indent     int
	width      int
	depth      int
	mode       interface{}
	wordList   map[string][2]string
}

//NewPprint construct
func NewPprint() *Pprint {
	pprint := new(Pprint)
	pprint.indent = 4
	temp := [2]string{"{", "}"}
	mapKey := keyword{"map", temp}
	temp = [2]string{"[", "]"}
	ArrayKey := keyword{"array", temp}
	SliceKey := keyword{"slice", temp}
	pprint.wordList = make(map[string][2]string)
	pprint.wordList["map"] = [2]string{"{", "}"}
	pprint.wordList["slice"] = [2]string{"[", "]"}
	pprint.wordList["array"] = [2]string{"[", "]"}
	pprint.formatList = append(pprint.formatList, mapKey)
	pprint.formatList = append(pprint.formatList, ArrayKey)
	pprint.formatList = append(pprint.formatList, SliceKey)
	return pprint
}

//pprint print list slice
func (p *Pprint) pprint(data interface{}) {
	p.mode = reflect.TypeOf(data).Kind()

	if p.mode == reflect.Map {
		p.PrintMaps(data)
	} else if p.mode == reflect.Slice {
		// fmt.Println("[slice]")
		// p.PrintSlices(data)
		p.PrintArrays(data)
	} else if p.mode == reflect.Array {
		// fmt.Println("[Array]")
		p.PrintArrays(data)
	} else {
		fmt.Println(data)
	}
	// Debug(t)
}

//test test
func test() {
	fmt.Println("test")
}

//PrintMaps print map
func (p *Pprint) PrintMaps(data interface{}) {
	// * Complete the chart
	// {'apple': 1,
	//  'banana': 3,
	//  'melon': 6,
	//  'orange': 2,
	//  'persimmon': 5,
	//  'pineapple': 4,
	//  'watermelon': 7}
}

//PrintArrays print map
func (p *Pprint) PrintArrays(data interface{}) {
	// * Complete the chart
	fmt.Println(p.wordList["array"][0])
	result := ChangeInterfaceArrayInterface(data)
	// fmt.Println(item)
	for _, item := range result {
		p.PrintIndent()
		fmt.Println(item)
	}
	fmt.Println(p.wordList["array"][1])
}

// PrintIndent print space indent size
func (p *Pprint) PrintIndent() {
	temp := strings.Repeat(" ", p.indent)
	fmt.Print(temp)
}

//ChangeInterfaceArrayInterface interface → Array
func ChangeInterfaceArrayInterface(data interface{}) []interface{} {
	s := reflect.ValueOf(data)
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret

}

//PrintSlices print slice
// func (p *Pprint) PrintSlices(data interface{}) {

// }

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
	// m := map[string]string{"foo": "0", "hello": "0"}
	// m := map[string]int{"foo": 0, "hello": 0}
	// s := []int{}
	b := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Debug(b[0])
	pprint := NewPprint()
	// pprint.pprint(m)
	// pprint.pprint(s)
	pprint.pprint(b)
	// n := map[int]string{1: "bar", 3: "world"}
	// n := [][]map[string]string
	// matrix := [][]string{{[]string{"0"}, []string{"0"}}}
	// Debug(reflect.TypeOf(m).Kind() == reflect.Map)
}
