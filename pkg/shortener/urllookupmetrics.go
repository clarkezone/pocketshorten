package shortener

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

type urlLookupMetrics struct {
	underlyingLoader urlLookupService
	entries          prometheus.Gauge
	lookups          *prometheus.CounterVec
}

func addMetrics(prefix string, l urlLookupService) (*urlLookupMetrics, urlLookupService) {
	clarkezoneLog.Debugf("add metrics called\n")
	mw := urlLookupMetrics{}
	mw.underlyingLoader = l

	mw.entries = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prefix + "_total_lookup_entries",
		Help: "Gauge containing number of url lookup entries stored",
	})

	mw.lookups = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: prefix + "_total_lookups",
		Help: "Counter containing number of key lookups by key",
	}, []string{"key"})
	metrics := &mw

	return &mw, metrics
}

func (lm *urlLookupMetrics) RecordStore() {
	lm.entries.Inc()
}

func (lm *urlLookupMetrics) Store(key string, en *URLEntry) error {
	return lm.underlyingLoader.Store(key, en)
}

func (lm *urlLookupMetrics) Lookup(key string) (*URLEntry, error) {
	lm.lookups.WithLabelValues(key).Inc()
	return lm.underlyingLoader.Lookup(key)
}

func (lm *urlLookupMetrics) Count() int {
	return lm.underlyingLoader.Count()
}

func (lm *urlLookupMetrics) Ready() bool {
	return lm.underlyingLoader.Ready()
}
