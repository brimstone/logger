package logger

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

type fakeWriter struct {
	length int64
	status int
	writer http.ResponseWriter
}

func (f *fakeWriter) Write(buf []byte) (int, error) {
	n, err := f.writer.Write(buf)
	f.length += int64(n)
	return n, err
}

func (f *fakeWriter) Header() http.Header {
	return f.writer.Header()
}

func (f *fakeWriter) WriteHeader(status int) {
	f.status = status
	f.writer.WriteHeader(status)
}

// Hijack implements http.Hijacker. If the underlying ResponseWriter is a
// Hijacker, its Hijack method is returned. Otherwise an error is returned.
// stolen from https://github.com/nytimes/gziphandler
func (f *fakeWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hj, ok := f.writer.(http.Hijacker); ok {
		return hj.Hijack()
	}
	return nil, nil, fmt.Errorf("http.Hijacker interface is not supported")
}

func HTTP(h http.Handler) http.Handler {
	log := New(&Options{})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &fakeWriter{
			writer: w,
		}

		host := r.Host
		h.ServeHTTP(wrapped, r)

		if wrapped.status == 0 {
			wrapped.status = 200
		}

		agent := ""
		if a, ok := r.Header["User-Agent"]; ok {
			agent = a[0]
		}
		log.Info("http request",
			log.Field("agent", agent),
			log.Field("ip", r.RemoteAddr),
			log.Field("length", wrapped.length),
			log.Field("method", r.Method),
			log.Field("proto", r.Proto),
			log.Field("status", wrapped.status),
			log.Field("uri", r.RequestURI),
			log.Field("host", host),
		)
	})
}
