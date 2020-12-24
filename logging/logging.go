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
	logging.formatTime = true
	logging.formatFunName = true
	logging.formatRunLine = true
	logging.formatList = make(map[string]string)
	logging.formatList["formatTime"] = "Run Time"
	logging.formatList["formatFunName"] = "Function Names"
	logging.formatList["formatRunLine"] = "Run Line number"
	logging.GetMaxLenghtForFormat()
	return logging
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
	l.FormatConfig()
	// l.test()
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
	for i := 0; i < 100; i++ {
		pc, _, line, ok := runtime.Caller(i)
		if ok {
			f := runtime.FuncForPC(pc)
			// fname := filepath.Base(file)
			fmt.Println(f.Name(), line)
			// fmt.Println(file)
			// fmt.Print(fname)
		} else {
			break
		}
		// fmt.Errorf(`check error. "%s" != "%s"`, a, b)
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
	logging := NewLogging()
	logging.Debug("debug test")
	// logging.Info("test")
	// log.Info("debug test")
	// log.GetLogger()
	// log.SetFlags(log.Llongfile)
	// log.Println("ログ1")
	// log.SetFlags(log.Lshortfile)
	// hello := fmt.Sprintf("'%-20s'\n", "test")
	// fmt.Println(hello)
}
