package main

import (
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
	logging.format = "%#v"
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
	return logging
}

//SetLevel set logging print level
func (l *Logging) SetLevel(level int) {
	l.levelNow = level
}

//SetFormat set logging print  format
func (l *Logging) SetFormat(format string) {
	l.format = format
}

//Debug row level print
func (l *Logging) Debug(text string) {
	if 0 >= l.levelNow {
		log.Printf(l.format, text)
	}
}

//Info row level print
func (l *Logging) Info(text string) {
	if 0 >= l.levelNow {
		log.Printf(l.format, text)
	}
}

//Warning row level print
func (l *Logging) Warning(text string) {
	if 0 >= l.levelNow {
		log.Printf(l.format, text)
	}
}

//Error row level print
func (l *Logging) Error(text string) {
	if 0 >= l.levelNow {
		log.Printf(l.format, text)
	}
}

//Critical row level print
func (l *Logging) Critical(text string) {
	if 0 >= l.levelNow {
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
