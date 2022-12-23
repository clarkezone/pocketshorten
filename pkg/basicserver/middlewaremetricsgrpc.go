package basicserver

import (
	"context"
	"fmt"
	"time"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// PromMetricsMiddlewareGrpc adds simple prometheus metrics type PromMetricsMiddlewareGrpc
type PromMetricsMiddlewareGrpc struct {
	opsProcessed    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}

// MetricsUnaryInterceptor is the interceptor that can be chained
func (l *PromMetricsMiddlewareGrpc) MetricsUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	result, err := handler(ctx, req)

	st, _ := status.FromError(err)

	httpDuration := time.Since(start)
	l.opsProcessed.WithLabelValues(fmt.Sprint(st.Code()), info.FullMethod).Inc()
	l.requestDuration.WithLabelValues(info.FullMethod).Observe(httpDuration.Seconds())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func newMiddlewareMetricsGrpc(prefix string) *PromMetricsMiddlewareGrpc {
	mw := PromMetricsMiddlewareGrpc{}
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

// NewPromMetricsMiddlewareGrpc constructs a new Logger middleware handler
func NewPromMetricsMiddlewareGrpc(prefix string) *PromMetricsMiddlewareGrpc {
	clarkezoneLog.Debugf("NewPromMetricsMiddlewareGrpc()")
	return newMiddlewareMetricsGrpc(prefix)
}
