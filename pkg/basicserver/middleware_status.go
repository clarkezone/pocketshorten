package basicserver

import (
	"fmt"
	"net/http"
)

// LoggingMiddlewareWeb adds logging functionality
type StatusMiddlewareWeb struct {
	handler      http.Handler
	statusWriter *StatusWriter
}

func (s *StatusMiddlewareWeb) Status() int {
	return s.statusWriter.status
}

type StatusWriter struct {
	responseWriter http.ResponseWriter
	status         int
}

func (s *StatusWriter) Status() int {
	return s.status
}

func (l *StatusWriter) Header() http.Header {
	return l.responseWriter.Header()
}

func (l *StatusWriter) Write(bytes []byte) (int, error) {
	return l.responseWriter.Write(bytes)
}

func (l *StatusWriter) WriteHeader(statusCode int) {
	fmt.Println("foo")
	l.status = statusCode
	l.responseWriter.WriteHeader(statusCode)
}

func (l *StatusMiddlewareWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("bar")
	l.statusWriter.responseWriter = w
	l.handler.ServeHTTP(l.statusWriter, r)
}

// NewLoggingMiddleware constructs a new Logger middleware handler
func NewStatusMiddlewareWeb(handlerToWrap http.Handler, sw *StatusWriter) *StatusMiddlewareWeb {
	return &StatusMiddlewareWeb{handlerToWrap, sw}
}

func NewStatusRecorder() *StatusWriter {
	return &StatusWriter{}
}
