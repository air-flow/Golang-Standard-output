package main

import (
	"fmt"
	"runtime"
	"time"
)

// Logging main struct
type Logging struct {
	format        string
	formatTime    bool
	formatFunName bool
	formatRunLine bool
	level         int
	DEBUG         int
	INFO          int
	WARRING       int
	ERROR         int
	CRITICAL      int
	levelList     [5]string
	formatList    map[string]string
	formatLenMax  int
	printMode     []string
	funcLog       []string
}

// NewLogging constructor
func NewLogging() *Logging {
	logging := new(Logging)
	logging.InitializeLevel()
	logging.InitializeLevelList()
	logging.InitializeFormatList()
	logging.InitializeFormatMode()
	logging.GetMaxLenghtForFormat()
	logging.level = logging.DEBUG
	logging.format = "%#v\n"
	return logging
}

//InitializeLevelList set list
func (l *Logging) InitializeLevelList() {
	l.levelList[l.DEBUG] = "debug"
	l.levelList[l.INFO] = "info"
	l.levelList[l.WARRING] = "warring"
	l.levelList[l.ERROR] = "error"
	l.levelList[l.CRITICAL] = "critical"
}

//InitializeLevel level set int
func (l *Logging) InitializeLevel() {
	l.DEBUG = 0
	l.INFO = 1
	l.WARRING = 2
	l.ERROR = 3
	l.CRITICAL = 4
}

//InitializeFormatList set map
func (l *Logging) InitializeFormatList() {
	l.formatList = make(map[string]string)
	l.formatList["formatTime"] = "Run Time"
	l.formatList["formatFunName"] = "Function Names"
	l.formatList["formatRunLine"] = "Run Line number"
}

//InitializeFormatMode format bool
func (l *Logging) InitializeFormatMode() {
	l.formatTime = true
	l.formatFunName = true
	l.formatRunLine = true
}

//InitializePrintMode print mode list string
func (l *Logging) InitializePrintMode() {
	l.printMode = append(l.printMode, "Ruby")
	l.Debug(l.printMode[0])
	// fmt.Println(reflect.TypeOf(l.printMode))
}

//GetMaxLenghtForFormat formatList get max string lenght
func (l *Logging) GetMaxLenghtForFormat() {
	var maxLenght int = 0
	for _, value := range l.formatList {
		if maxLenght < len(value) {
			maxLenght = len(value)
		}
	}
	l.formatLenMax = maxLenght
}

//GetLogger Field Variable print
func (l *Logging) GetLogger() {
	fmt.Printf("format => %q\n", l.format)
	fmt.Printf("format => %d\n", l.level)
}

//SetLevel set logging print level
func (l *Logging) SetLevel(level int) {
	// !VALIDATION CHECK
	l.level = level
}

//SetFormat set logging print  format
func (l *Logging) SetFormat(format string) {
	// !VALIDATION CHECK
	l.format = format
}

// Printf loggging
func (l *Logging) Printf(text interface{}) {
	// l.test()
	l.FormatConfig()
	fmt.Printf(l.format, text)
}

//FormatConfig check status
func (l *Logging) FormatConfig() {
	var formatTemplate = fmt.Sprintf("-%ds", l.formatLenMax+1)
	if l.formatTime {
		var temp = fmt.Sprintf("%"+formatTemplate, l.formatList["formatTime"])
		fmt.Printf(temp+"=> %s\n", time.Now())
	}
	if l.formatFunName {
		pc, _, _, _ := runtime.Caller(3)
		f := runtime.FuncForPC(pc)
		var temp = fmt.Sprintf("%"+formatTemplate, l.formatList["formatFunName"])
		fmt.Printf(temp+"=> %s\n", f.Name())
	}
	if l.formatRunLine {
		var temp = fmt.Sprintf("%"+formatTemplate, l.formatList["formatRunLine"])
		_, _, line, _ := runtime.Caller(3)
		fmt.Printf(temp+"=> %d\n", line)
	}
}

func (l *Logging) test() {
	if false {
		return
	}
	fmt.Println("test return")
}

//Debug row level print
func (l *Logging) Debug(text interface{}) {
	if l.DEBUG >= l.level {
		l.Printf(text)
	}
}

//Info row level print
func (l *Logging) Info(text interface{}) {
	if l.INFO >= l.level {
		l.Printf(text)
	}
}

//Warning row level print
func (l *Logging) Warning(text interface{}) {
	if l.WARRING >= l.level {
		l.Printf(text)
	}
}

//Error row level print
func (l *Logging) Error(text interface{}) {
	if l.ERROR >= l.level {
		l.Printf(text)
	}
}

//Critical row level print
func (l *Logging) Critical(text interface{}) {
	if l.CRITICAL >= l.level {
		l.Printf(text)
	}
}

//checkMain
func checkMain() {
	logging := NewLogging()
	logging.Debug(1)
	logging.SetLevel(logging.INFO)
	logging.Info("test")
}

func main() {
	checkMain()
	// logging := NewLogging()
	// logging.InitializePrintMode()
}
