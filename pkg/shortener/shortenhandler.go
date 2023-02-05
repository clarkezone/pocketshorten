package shortener

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

type storeLoader interface {
	Init(urlLookupService) error
}

type URLEntry struct {
	ShortLink       string
	DestinationLink string
	LinkGroup       string
	Created         time.Time
}

type urlLookupService interface {
	Store(string, *URLEntry) error
	Lookup(string) (*URLEntry, error)
	Count() int
	Ready() bool
}

// NewDictLookupHandler creates a new instance of type
//
//lint:ignore U1000 reason backend not selected
func NewDictLookupHandler() *ShortenHandler {
	vl := &viperLoader{}
	ds := newDictStore(vl)
	lh := &ShortenHandler{storage: ds}
	return lh
}

// NewGrpcLookupHandler returns a new lookuphandler instance
func NewGrpcLookupHandler(s string) (*ShortenHandler, error) {
	// dictstore
	// grpcloader
	ds, err := newGrpcStore(s)
	if err != nil {
		return nil, err
	}
	lh := &ShortenHandler{storage: ds}
	return lh, nil
}

// ShortenHandler core logic
type ShortenHandler struct {
	storage urlLookupService
}

// RegisterHandlers attaches handlers to Mux that is passed in
func (lh *ShortenHandler) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", lh.redirectHandler)
	mux.HandleFunc("/ready", lh.readyHandler)
	mux.HandleFunc("/live", lh.liveHandler)
}

func (lh *ShortenHandler) redirectHandler(w http.ResponseWriter, r *http.Request) {
	if !lh.storage.Ready() {
		writeOutputError(w, "server error: not configured", http.StatusInternalServerError)
		return
	}
	//requested := r.URL.Query().Get("shortlink")
	requested, err := sanitize(r.URL.Path)

	// TODO update scalbility tests
	clarkezoneLog.Debugf("path: :%v: sanitized:%v:", requested)
	if err != nil {
		writeOutputError(w, fmt.Sprintf("input sanitization failed: unabled to process request %v", err), http.StatusBadRequest)
		return
	}

	if requested == "" {
		writeOutputError(w, "please supply shortlink query parameter", http.StatusNotFound)
		return
	}
	uri, err := lh.storage.Lookup(requested)
	if err != nil {
		writeOutputError(w, fmt.Sprintf("shortlink %v notfound", requested), http.StatusNotFound)
		return
	}
	clarkezoneLog.Debugf("redirecting to %v", uri)

	http.Redirect(w, r, uri.DestinationLink, http.StatusMovedPermanently)
}

func (lh *ShortenHandler) liveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (lh *ShortenHandler) readyHandler(w http.ResponseWriter, r *http.Request) {
	if !lh.storage.Ready() {
		writeOutputError(w, "Service not available", http.StatusServiceUnavailable)
	}
	w.WriteHeader(http.StatusOK)
}

func writeOutputError(w http.ResponseWriter, message string, code int) {
	clarkezoneLog.Debugf("Error reported to user %v", message)
	http.Error(w, message, code)
}

func sanitize(input string) (string, error) {
	const maxinput int = 20
	sa := strings.TrimLeft(input, "/")
	clarkezoneLog.Debugf("sanitized path: %v", sa)
	le := len(sa)
	if len(sa) > maxinput {
		return "", fmt.Errorf("bad input expected < %v chars received %v chars", maxinput, le)
	}
	return sa, nil
}
