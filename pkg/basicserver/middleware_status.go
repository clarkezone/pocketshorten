package basicserver

import (
	"fmt"
	"net/http"
)

// StatusWriter tracks status codes for responses
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
	if l.responseWriter != nil {
		l.responseWriter.WriteHeader(statusCode)
	}
}

// NewStatusRecorder constructs a new Logger middleware handler
func NewStatusRecorder(w http.ResponseWriter) *StatusWriter {
	sw := &StatusWriter{}
	sw.responseWriter = w
	return sw
}
