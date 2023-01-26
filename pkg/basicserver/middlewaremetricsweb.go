package basicserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

// PromMetricsMiddlewareWeb adds simple prometheus metrics type PromMetricsMiddlewareWeb
type PromMetricsMiddlewareWeb struct {
	handler         http.Handler
	opsProcessed    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
	statusGetter    StatusRecorder
}

func (l *PromMetricsMiddlewareWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	httpDuration := time.Since(start)
	l.opsProcessed.WithLabelValues(fmt.Sprint(l.statusGetter.Status()), r.Method).Inc()
	l.requestDuration.WithLabelValues(r.RequestURI).Observe(httpDuration.Seconds())
}

func newMiddlewareMetricsWeb(handlerToWrap http.Handler, prefix string, sg StatusRecorder) *PromMetricsMiddlewareWeb {
	mw := PromMetricsMiddlewareWeb{}
	mw.statusGetter = sg
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

// NewPromMetricsMiddlewareWeb constructs a new Logger middleware handler
func NewPromMetricsMiddlewareWeb(prefix string, handlerToWrap http.Handler, status StatusRecorder) *PromMetricsMiddlewareWeb {
	clarkezoneLog.Debugf("NewPromMetricsMiddleware()")
	return newMiddlewareMetricsWeb(handlerToWrap, prefix, status)
}
