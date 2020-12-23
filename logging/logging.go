package main

import (
	"log"
)

// Logging main struct
type Logging struct {
	format string
	level  int
}

// NewLogging constructor
func NewLogging() *Logging {
	logging := new(Logging)
	logging.format = "%#v"
	logging.level = 0
	return logging
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
	if 0 >= l.level {
		log.Printf(l.format, text)
	}
}

//Info row level print
func (l *Logging) Info(text string) {
	if 0 >= l.level {
		log.Printf(l.format, text)
	}
}

//Warning row level print
func (l *Logging) Warning(text string) {
	if 0 >= l.level {
		log.Printf(l.format, text)
	}
}

//Error row level print
func (l *Logging) Error(text string) {
	if 0 >= l.level {
		log.Printf(l.format, text)
	}
}

//Critical row level print
func (l *Logging) Critical(text string) {
	if 0 >= l.level {
		log.Printf(l.format, text)
	}
}

func main() {
	// log.Printf(format, test)
	log := NewLogging()
	log.Debug("test")
	log.SetLevel(1)
	log.Debug("test")
}
