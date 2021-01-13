package main

import "fmt"

func main() {
	// var str interface{}
	// str = []int{0, 1, 2, 3, 4, 5, 6}
	// var list []interface{}
	// ret := make([]interface{}, str.Len)
	// list = []int{0, 1, 2, 3, 4, 5, 6}
	// fmt.Println(str)
	// fmt.Println(list)
	mapOfSlices := map[string][]string{
		"first":  {},
		"second": []string{"one", "two", "three", "four", "five"},
		"third":  []string{"quarter", "half"},
	}
	fmt.Println(mapOfSlices["third"][0])
	wordList := make(map[string][2]string)
	wordList["map"] = [2]string{"{", "}"}
	wordList["list"] = [2]string{"[", "]"}
	wordList["array"] = [2]string{"[", "]"}
	fmt.Println(wordList)
}
