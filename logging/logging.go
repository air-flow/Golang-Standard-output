package main

import "log"

// Logging main struct
type Logging struct {
	format string
	level  int
}

//setLevel set logging print level
func (l *Logging) setLevel() {

}

//debug row level print
func (l *Logging) debug() {

}

//setFormat set logging print  format
func (l *Logging) setFormats() {
	l.format = ""
}

func main() {
	test := "test"
	format := "%#v"
	log.Printf(format, test)

}
