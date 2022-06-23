package logger

import (
	"fmt"
	"runtime"
)

func (l *Logger) Recover(f func(message, method, filename string, linenumber int64)) {
	r := recover()
	if r == nil {
		return
	}
	pc, filename, linenumber, _ := runtime.Caller(3)
	f(fmt.Sprintf("%s", r), runtime.FuncForPC(pc).Name(), filename, int64(linenumber))
}
