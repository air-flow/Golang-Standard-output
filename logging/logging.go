package main

import (
	"fmt"
	"log"
	"runtime"
)

// Logging main struct
type Logging struct {
	format        string
	formatTime    bool
	formatFunName bool
	level         int
	DEBUG         int
	INFO          int
	WARRING       int
	ERROR         int
	CRITICAL      int
	levelList     [5]string
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
	logging.formatTime = false
	logging.formatFunName = false
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

// Printf loggging
func (l *Logging) Printf(text string) {
	log.Printf(l.format, text)
}

//Debug row level print
func (l *Logging) Debug(text string) {
	if l.DEBUG >= l.level {
		l.Printf(text)
	}
}

//Info row level print
func (l *Logging) Info(text string) {
	if l.INFO >= l.level {
		l.Printf(text)
	}
}

//Warning row level print
func (l *Logging) Warning(text string) {
	if l.WARRING >= l.level {
		l.Printf(text)
	}
}

//Error row level print
func (l *Logging) Error(text string) {
	if l.ERROR >= l.level {
		l.Printf(text)
	}
}

//Critical row level print
func (l *Logging) Critical(text string) {
	if l.CRITICAL >= l.level {
		l.Printf(text)
	}
}

func Log(data string) {

	pc, file, line, _ := runtime.Caller(1)
	fmt.Println(pc)
	f := runtime.FuncForPC(pc)
	log.Printf("\ncall:%s\ndata:%s\nfile:%s:%d\n", f.Name(), data, file, line)
}

func main() {
	// log.Printf(format, test)
	log := NewLogging()
	// log.Debug("debug test")
	log.SetLevel(1)
	// log.Info("debug test")
	// log.GetLogger()
}
