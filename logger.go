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

	counter       map[string]*int64
	file          bool
	gauge         map[string]interface{}
	method        string
	metricPrinter bool
	metricsDelay  time.Duration
}

type Options struct {
	Method string
	File   bool
	Delay  time.Duration
}

func New(o ...*Options) *Logger {
	l := &Logger{
		logrus: logrus.New(),
	}
	for _, option := range o {
		if option.Method != "" {
			l.method = option.Method
		}
		l.file = option.File
		l.metricsDelay = time.Second * 60
		if option.Delay != 0 {
			l.metricsDelay = option.Delay
		}
	}
	l.logrus.SetLevel(logrus.DebugLevel)
	l.logrus.Out = os.Stderr
	//l.logrus.SetFormatter(&logrus.TextFormatter{})
	l.gauge = make(map[string]interface{})
	l.counter = make(map[string]*int64)
	return l
}

func Method(method string, o ...*Options) *Logger {
	var options []*Options
	options = append(options,
		&Options{
			Method: method,
			File:   true,
		})
	options = append(options, o...)

	return New(options...)
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
	if l.method != "" {
		e = e.WithField("method", l.method)
	}
	if l.file {
		_, filename, linenumber, ok := runtime.Caller(1)
		if !ok {
			panic("How did you get here?")
		}
		e = e.WithField("file", filename+":"+strconv.FormatInt(int64(linenumber), 10))
	}
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
