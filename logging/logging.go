package main

import (
	"fmt"
	"log"
)

// Logging main struct
type Logging struct {
	format    string
	level     int
	DEBUG     int
	INFO      int
	WARRING   int
	ERROR     int
	CRITICAL  int
	levelList [5]string
}

// NewLogging constructor
func NewLogging() *Logging {
	logging := new(Logging)
	logging.DEBUG = 0
	logging.INFO = 1
	logging.WARRING = 2
	logging.ERROR = 3
	logging.CRITICAL = 4
	logging.levelList[logging.DEBUG] = "debug"
	logging.levelList[logging.INFO] = "info"
	logging.levelList[logging.WARRING] = "warring"
	logging.levelList[logging.ERROR] = "error"
	logging.levelList[logging.CRITICAL] = "critical"
	logging.level = logging.DEBUG
	logging.format = "%#v"
	return logging
}

//GetLogger Field Variable print
func (l *Logging) GetLogger() {
	fmt.Printf("format => %q\n", l.format)
	fmt.Printf("format => %d\n", l.level)
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
	if l.DEBUG >= l.level {
		log.Printf(l.format, text)
	}
}

//Info row level print
func (l *Logging) Info(text string) {
	if l.INFO >= l.level {
		log.Printf(l.format, text)
	}
}

//Warning row level print
func (l *Logging) Warning(text string) {
	if l.WARRING >= l.level {
		log.Printf(l.format, text)
	}
}

//Error row level print
func (l *Logging) Error(text string) {
	if l.ERROR >= l.level {
		log.Printf(l.format, text)
	}
}

//Critical row level print
func (l *Logging) Critical(text string) {
	if l.CRITICAL >= l.level {
		log.Printf(l.format, text)
	}
}

func main() {
	// log.Printf(format, test)
	log := NewLogging()
	log.Debug("debug test")
	log.SetLevel(1)
	log.Info("debug test")
	log.GetLogger()
}
