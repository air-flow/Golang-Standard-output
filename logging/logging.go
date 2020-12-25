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
	printMode     string
	funcLog       [10]string
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
	l.formatList["formatRunLine"] = "Run Line number"s
}

//InitializeFormatMode format bool
func (l *Logging) InitializeFormatMode()  {
	logging.formatTime = true
	logging.formatFunName = true
	logging.formatRunLine = true
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
	l.level = level
}

//SetFormat set logging print  format
func (l *Logging) SetFormat(format string) {
	l.format = format
}

// Printf loggging
func (l *Logging) Printf(text string) {
	// l.test()
	l.FormatConfig()
	fmt.Printf(l.format, text)
}

//FormatConfig check status
func (l *Logging) FormatConfig() {
	var formatTemplate = fmt.Sprintf("-%ds", l.formatLenMax+1)
	// fmt.Printf("%"+formatTemplate+" ", "temp")
	if l.formatTime {
		var temp = fmt.Sprintf("%"+formatTemplate, l.formatList["formatTime"])
		fmt.Printf(temp+"=> %s\n", time.Now())
	}
	if l.formatFunName {
		pc, _, _, _ := runtime.Caller(3)
		f := runtime.FuncForPC(pc)
		var temp = fmt.Sprintf("%"+formatTemplate, l.formatList["formatFunName"])
		fmt.Printf(temp+"=> %s\n", f.Name())
		// file, line := f.FileLine(f.Entry())
		// fmt.Println(file, line)
		// println(f.Name())
		// println("test")
	}
	if l.formatRunLine {
		var temp = fmt.Sprintf("%"+formatTemplate, l.formatList["formatRunLine"])
		_, _, line, _ := runtime.Caller(3)
		// f := runtime.FuncForPC(pc)
		fmt.Printf(temp+"=> %d\n", line)
		// fmt.Println(f.Name(), line)
	}
}

func (l *Logging) test() {
	if false {
		return
	}
	fmt.Println("test return")
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

//checkMain
func checkMain() {
	logging := NewLogging()
	logging.SetLevel(logging.INFO)
	logging.Debug("debug test")
	logging.Info("test")
}

func main() {
	logging := NewLogging()
	logging.test()
}
