package shortener

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type urlLookupMetrics struct {
	underlyingLoader urlLookupService
	entries          prometheus.Counter
	lookups          *prometheus.CounterVec
}

func addMetrics(prefix string, l urlLookupService) urlLookupService {
	mw := urlLookupMetrics{}
	mw.underlyingLoader = l

	mw.entries = promauto.NewCounter(prometheus.CounterOpts{
		Name: prefix + "_total_lookup_entries",
		Help: "Counter containing number of url lookup entries stored",
	})

	mw.lookups = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: prefix + "_total_lookups",
		Help: "Counter containing number of key lookups by key",
	}, []string{"key"})
	metrics := &mw
	return metrics
}

func (lm *urlLookupMetrics) Store(key string, en *URLEntry) error {
	lm.entries.Inc()
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
