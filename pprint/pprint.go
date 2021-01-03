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
	width      int
	depth      int
}

//NewPprint construct
func NewPprint() *Pprint {
	pprint := new(Pprint)
	mapKey := keyword{"map", "{", "}"}
	ArrayKey := keyword{"array", "[", "]"}
	SliceKey := keyword{"slice", "[", "]"}
	pprint.formatList = append(pprint.formatList, mapKey)
	pprint.formatList = append(pprint.formatList, ArrayKey)
	pprint.formatList = append(pprint.formatList, SliceKey)
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
	//! interface{} Convert to array
	// * Complete the chart
	// ['Tokyo',
	//  'Osaka',
	//  'Nagoya',
	//  'Fukuoka',
	//  'Hokkaido',
	//  'Kobe',
	//  'Yokohama',
	//  'Okinawa']
	// for _, v := range data.([10]int) {
	// 	fmt.Println(v)
	// }
	fmt.Println(len(data.(string)))
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
	// m := map[string]string{"foo": "0", "hello": "0"}
	// m := map[string]int{"foo": 0, "hello": 0}
	// s := []int{}
	b := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	pprint := NewPprint()
	// pprint.pprint(m)
	// pprint.pprint(s)
	pprint.pprint(b)

	// n := map[int]string{1: "bar", 3: "world"}
	// n := [][]map[string]string
	// matrix := [][]string{{[]string{"0"}, []string{"0"}}}
	// Debug(reflect.TypeOf(m).Kind() == reflect.Map)
}
