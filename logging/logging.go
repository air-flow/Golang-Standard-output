package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
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
	logging.format = "%#v\n"
	logging.formatTime = false
	logging.formatFunName = true
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
	l.FormatConfig()
	fmt.Printf(l.format, text)
}

//FormatConfig check status
func (l *Logging) FormatConfig() {
	if l.formatTime {
		fmt.Printf("Run Time => %s\n", time.Now())
	}
	if l.formatFunName {
		pc, _, _, _ := runtime.Caller(3)
		f := runtime.FuncForPC(pc)
		fmt.Printf("Function Name => %s\n", strings.Split(f.Name(), ".")[0])
		file, line := f.FileLine(f.Entry())
		fmt.Println(file, line)
		// println(f.Name())
		// println("test")
	}
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

func main() {
	// log.Printf(format, test)
	// logging := NewLogging()
	// logging.Debug("debug test")
	// logging.Info("test")
	// log.Info("debug test")
	// log.GetLogger()
	log.SetFlags(log.Llongfile)
	log.Println("ログ1")
	log.SetFlags(log.Lshortfile)
}
