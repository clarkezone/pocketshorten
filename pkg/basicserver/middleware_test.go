package basicserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_middleware(t *testing.T) {
	mux := DefaultMux()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %v", r.URL)
		w.WriteHeader(http.StatusOK)
	})
	var wrappedmux http.Handler
	wrappedmux = NewLoggingMiddleware(mux)
	sg := NewStatusRecorder()
	statusMw := NewStatusMiddlewareWeb(wrappedmux, sg)
	wrappedmux = statusMw

	s := httptest.NewServer(wrappedmux)
	defer s.Close()

	resp, err := http.DefaultClient.Get(s.URL + "/test")
	if err != nil {
		t.Fatalf("Error")
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected response")
	}

	if statusMw.Status() != 200 {
		t.Fatalf("Statusmiddleware didn't work.  Expected 200 received %v", statusMw.Status())
	}

}
