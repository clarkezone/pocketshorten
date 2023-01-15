package shortener

import (
	"fmt"
	"net/http"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

type storeLoader interface {
	Init(urlLookupService)
}

type urlLookupService interface {
	Store(string, string) error
	Lookup(string) (string, error)
}

//lint:ignore U1000 reason backend not selected
func newDictLookupHandler() *ShortenHandler {
	// need viperStoreLoader
	ds := NewDictStore(nil)
	lh := &ShortenHandler{storage: ds}
	return lh
}

func NewGrpcLookupHandler(s string) (*ShortenHandler, error) {
	// dictstore
	// grpcloader
	ds, err := NewGrpcStore(s)
	if err != nil {
		return nil, err
	}
	lh := &ShortenHandler{storage: ds}
	return lh, nil
}

type ShortenHandler struct {
	storage urlLookupService
}

func (lh *ShortenHandler) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", lh.redirectHandler)
}

func (lh *ShortenHandler) redirectHandler(w http.ResponseWriter, r *http.Request) {
	//TODO telemetry
	requested := r.URL.Query().Get("shortlink")

	if requested == "" {
		clarkezoneLog.Errorf("redirecthandler called with  missingshortlink")
		writeOutputError(w, "please supply shortlink query parameter")
		return
	}
	uri, err := lh.storage.Lookup(requested)
	if err != nil {
		writeOutputError(w, fmt.Sprintf("shortlink %v notfound", requested))
		return
	}
	clarkezoneLog.Debugf("redirecting to %v", uri)

	http.Redirect(w, r, uri, http.StatusMovedPermanently)

}

func writeOutputError(w http.ResponseWriter, message string) {
	clarkezoneLog.Debugf("Error reported to user %v", message)
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte(message))
	if err != nil {
		clarkezoneLog.Debugf("writeOutputError: Failed to write bytes %v\n", err)
	}
}
