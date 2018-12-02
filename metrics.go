package logger

import (
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
)

func (l *Logger) printMetrics() {
	log := New(&Options{})
	for {
		e := logrus.NewEntry(log.logrus)
		for k, v := range l.gauge {
			e = e.WithField(k, v)
		}
		for k, v := range l.counter {
			e = e.WithField(k, *v)
		}
		e.Debug("metrics")
		time.Sleep(l.metricsDelay)
	}
}

func (l *Logger) Gauge(key string, value interface{}) {
	l.gauge[key] = value
	if !l.metricPrinter && l.metricsDelay > 0 {
		l.metricPrinter = true
		go l.printMetrics()
	}
}

func (l *Logger) Counter(key string, increment int64) {
	if _, ok := l.counter[key]; !ok {
		l.counter[key] = new(int64)
	}
	atomic.AddInt64(l.counter[key], increment)
	if !l.metricPrinter && l.metricsDelay > 0 {
		l.metricPrinter = true
		go l.printMetrics()
	}
	time.Sleep(0)
}

func (l *Logger) RemoveGauge(key string) {
	delete(l.gauge, key)
}
