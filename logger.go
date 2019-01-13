package logger

import (
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

type Fields logrus.Fields

type Logger struct {
	logrus *logrus.Logger

	counter       map[string]*int64
	file          bool
	gauge         map[string]interface{}
	metricPrinter bool
	metricsDelay  time.Duration
}

type Options struct {
	File  bool
	Delay time.Duration
}

func New(o ...*Options) *Logger {
	l := &Logger{
		logrus: logrus.New(),
	}
	for _, option := range o {
		l.file = option.File
		l.metricsDelay = time.Second * 60
		if option.Delay != 0 {
			l.metricsDelay = option.Delay
		}
	}
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if logLevel == "warn" {
		l.logrus.SetLevel(logrus.WarnLevel)
	} else if logLevel == "info" {
		l.logrus.SetLevel(logrus.InfoLevel)
	} else if logLevel == "error" {
		l.logrus.SetLevel(logrus.ErrorLevel)
	} else {
		l.logrus.SetLevel(logrus.DebugLevel)
	}
	l.logrus.Out = os.Stderr
	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		l.logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		l.logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	l.gauge = make(map[string]interface{})
	l.counter = make(map[string]*int64)
	return l
}

func (l *Logger) Profile(then time.Time) {
	l.prepFields(1, l.Field("duration", time.Now().Sub(then))).Debug("Profile")
}

type FieldPair struct {
	key   string
	value interface{}
}

func (l *Logger) Field(key string, value interface{}) FieldPair {
	return FieldPair{
		key:   key,
		value: value,
	}
}

func (l *Logger) prepFields(caller int, fps ...FieldPair) *logrus.Entry {
	e := logrus.NewEntry(l.logrus)
	pc, filename, linenumber, ok := runtime.Caller(caller + 2)
	if !ok {
		panic("How did you get here?")
	}
	if l.file {
		e = e.WithField("file", filename+":"+strconv.FormatInt(int64(linenumber), 10))
	}
	e = e.WithField("method", runtime.FuncForPC(pc).Name())
	for _, fp := range fps {
		e = e.WithField(fp.key, fp.value)
	}
	return e
}

func (l *Logger) Debug(msg string, fps ...FieldPair) {
	l.prepFields(0, fps...).Debug(msg)
}

func (l *Logger) Info(msg string, fps ...FieldPair) {
	l.prepFields(0, fps...).Info(msg)
}

func (l *Logger) Error(msg string, fps ...FieldPair) {
	l.prepFields(0, fps...).Error(msg)
}

func (l *Logger) Warn(msg string, fps ...FieldPair) {
	l.prepFields(0, fps...).Warn(msg)
}
