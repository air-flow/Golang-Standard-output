package main

import (
	"fmt"
	"log"
)

// Logging main struct
type Logging struct {
	format string
	level  int
}

//SetLevel set logging print level
func (l *Logging) SetLevel(level int) {
	l.level = level
}

//SetFormat set logging print  format
func (l *Logging) SetFormat(format string) {
	l.format = format
}

//Debug row level print
func (l *Logging) Debug(text string) {
	if 0 > l.level {
		fmt.Printf(l.format, text)
	}
}
func main() {
	test := "test"
	format := "%#v"
	log.Printf(format, test)
}
