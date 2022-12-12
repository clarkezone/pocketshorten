package basicserver

import "net/http"

type statusRecorder struct {
	http.ResponseWriter
	code int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.code = code
	rec.ResponseWriter.WriteHeader(code)
}
