package logger

import (
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type Fields logrus.Fields

type Logger struct {
	logrus *logrus.Logger
	method string
}

type Options struct {
	Method string
}

func New(o *Options) *Logger {
	l := &Logger{
		logrus: logrus.New(),
	}
	if o != nil {
		if o.Method != "" {
			l.method = o.Method
		}
	}
	l.logrus.SetLevel(logrus.DebugLevel)
	l.logrus.Out = os.Stderr
	//l.logrus.SetFormatter(&logrus.TextFormatter{})
	return l
}

func Method(method string) *Logger {
	return New(&Options{
		Method: method,
	})
}

func (l *Logger) Profile(then time.Time) {
	l.Debug("Profile",
		l.Field("duration", time.Now().Sub(then)))
}

type FieldPair struct {
	key   string
	value interface{}
}

func (l *Logger) Println(msg string) {
	l.Debug(msg)
}

func (l *Logger) Field(key string, value interface{}) FieldPair {
	return FieldPair{
		key:   key,
		value: value,
	}
}

func (l *Logger) Debug(msg string, fps ...FieldPair) {
	e := logrus.NewEntry(l.logrus)
	e = e.WithField("method", l.method)
	_, filename, linenumber, ok := runtime.Caller(1)
	if !ok {
		panic("How did you get here?")
	}
	e = e.WithField("file", filename+":"+strconv.FormatInt(int64(linenumber), 10))
	for _, fp := range fps {
		e = e.WithField(fp.key, fp.value)
	}
	e.Debug(msg)
}

func (l *Logger) Info(msg string, fps ...FieldPair) {
	e := logrus.NewEntry(l.logrus)
	/*
		e = e.WithField("method", l.method)
		_, filename, linenumber, ok := runtime.Caller(1)
		if !ok {
			panic("How did you get here?")
		}
		e = e.WithField("file", filename+":"+strconv.FormatInt(int64(linenumber), 10))
	*/
	for _, fp := range fps {
		e = e.WithField(fp.key, fp.value)
	}
	e.Info(msg)
}
