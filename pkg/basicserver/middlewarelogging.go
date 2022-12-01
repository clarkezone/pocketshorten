package basicserver

import (
	"net/http"
	"time"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

// LoggingMiddleware adds logging functionality
type LoggingMiddleware struct {
	handler http.Handler
}

func (l *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	clarkezoneLog.Debugf("LogMW: %s %s %v", r.Method, r.URL, time.Since(start))
}

// NewLoggingMiddleware constructs a new Logger middleware handler
func NewLoggingMiddleware(handlerToWrap http.Handler) *LoggingMiddleware {
	return &LoggingMiddleware{handlerToWrap}
}
