package logger

import (
	"sync/atomic"
	"time"
)

func (l *Logger) printMetrics() {
	log := New(&Options{})
	for {
		/*
			e := logrus.NewEntry(log.logrus)
			for k, v := range l.gauge {
				e = e.WithField(k, v)
			}
			e.Debug("metrics")
		*/
		for k, v := range l.counters {
			if v.lastObserved.IsZero() {
				v.lastObserved = time.Now()
				continue
			}
			timeFrame := int64((time.Now().Sub(v.lastObserved) / time.Second))
			log.Info("metric",
				log.Field("name", k),
				log.Field("current", *v.value),
				log.Field("rps", (*v.value-v.lastValue)/timeFrame),
			)
			v.lastValue = *v.value
			v.lastObserved = time.Now()
		}
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
	if _, ok := l.counters[key]; !ok {
		l.counters[key] = &counter{value: new(int64)}
	}
	atomic.AddInt64(l.counters[key].value, increment)
	if !l.metricPrinter && l.metricsDelay > 0 {
		l.metricPrinter = true
		go l.printMetrics()
	}
	time.Sleep(0)
}

func (l *Logger) RemoveGauge(key string) {
	delete(l.gauge, key)
}
