package basicserver

import (
	"fmt"
	"net/http"
	"time"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type statusRecorder struct {
	http.ResponseWriter
	code int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.code = code
	rec.ResponseWriter.WriteHeader(code)
}

// PromMetricsMiddleware adds simple prometheus metrics type PromMetricsMiddleware
type PromMetricsMiddleware struct {
	handler         http.Handler
	opsProcessed    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}

func (l *PromMetricsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := statusRecorder{w, 200}
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	httpDuration := time.Since(start)
	l.opsProcessed.WithLabelValues(fmt.Sprint(rec.code), r.Method).Inc()
	l.requestDuration.WithLabelValues(r.RequestURI).Observe(httpDuration.Seconds())
}

func newMiddleware(handlerToWrap http.Handler, prefix string) *PromMetricsMiddleware {
	mw := PromMetricsMiddleware{}
	mw.handler = handlerToWrap
	mw.opsProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: prefix + "_totalops",
		Help: "The total number of processed http requests for testserver",
	}, []string{"responsecode", "method"})
	mw.requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    prefix + "_duration_seconds",
		Help:    "Histogram of duration in seconds",
		Buckets: []float64{1, 2, 5, 7, 10},
	},
		[]string{"endpoint"})
	prometheus.MustRegister(mw.requestDuration)
	return &mw
}

// NewPromMetricsMiddleware constructs a new Logger middleware handler
func NewPromMetricsMiddleware(prefix string, handlerToWrap http.Handler) *PromMetricsMiddleware {
	clarkezoneLog.Debugf("NewPromMetricsMiddleware()")
	return newMiddleware(handlerToWrap, prefix)
}
