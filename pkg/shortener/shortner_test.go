package shortener

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/clarkezone/pocketshorten/internal"

	"github.com/clarkezone/pocketshorten/pkg/basicserver"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

// TestMain initizlie all tests
func TestMain(m *testing.M) {
	clarkezoneLog.Init(logrus.DebugLevel)
	internal.SetupGitRoot()
	code := m.Run()
	os.Exit(code)
}

type testLookupHandler struct {
	lookup urlLookupService
}

func (tlh *testLookupHandler) Init(ls urlLookupService) error {
	tlh.lookup = ls
	ue1 := URLEntry{"one", "two", "three", time.Now()}
	ue2 := URLEntry{"four", "three", "one", time.Now()}
	err := ls.Store("one", &ue1)
	err2 := ls.Store("three", &ue2)
	if err != nil || err2 != nil {
		return err
	}
	return nil
}

func Test_dictStore(t *testing.T) {
	// pass in viper loader
	tlh := &testLookupHandler{}
	handler := newDictStore(tlh)
	if handler == nil {
		t.Errorf("handler is nil")
	}
	val, err := handler.Lookup("one")
	if err != nil {
		t.Errorf("lookup failed")
	}
	if val.DestinationLink != "two" {
		t.Errorf("Wrong value")
	}
}

func Test_dictLookupHandler(t *testing.T) {
	handler := NewDictLookupHandler("")
	if handler == nil {
		t.Errorf("handler is nil")
	}
}

func Test_viperlookuphandlerinit(t *testing.T) {
	initviperconfig(t)

	handler := NewDictLookupHandler("")
	if handler != nil {

		//lint:ignore SA5011 reason test
		if handler.storage.Count() != 6 {
			t.Errorf("wrong number of items in storage")
		}
	}
}

func Test_viperlookuphandlerlookupbadurl(t *testing.T) {
	initviperconfig(t)

	handler := NewDictLookupHandler("")
	if handler == nil {
		t.Errorf("handler is nil")
	}

	req, err := http.NewRequest("GET", "/one", nil)
	if err != nil {
		t.Errorf("error creating request")
	}

	rr := httptest.NewRecorder()
	handler.redirectHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("wrong status code")
	}
}

func Test_viperrejectlonginput(t *testing.T) {
	initviperconfig(t)

	handler := NewDictLookupHandler("")
	if handler == nil {
		t.Errorf("handler is nil")
	}

	req, err := http.NewRequest("GET", "/onedddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd", nil)
	if err != nil {
		t.Errorf("error creating request")
	}

	rr := httptest.NewRecorder()
	handler.redirectHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("wrong status code")
	}
}

func Test_viperignorequery(t *testing.T) {
	initviperconfig(t)

	handler := NewDictLookupHandler("")
	if handler == nil {
		t.Errorf("handler is nil")
	}

	req, err := http.NewRequest("GET", "/key1?ignorethis", nil)
	if err != nil {
		t.Errorf("error creating request")
	}

	rr := httptest.NewRecorder()
	handler.redirectHandler(rr, req)

	if rr.Code != http.StatusMovedPermanently {
		t.Errorf("wrong status code result %v", rr.Code)
	}
}

func Test_viperlookuphandlergoodurlmissingkey(t *testing.T) {
	initviperconfig(t)

	handler := NewDictLookupHandler("")
	if handler == nil {
		t.Errorf("handler is nil")
	}

	req, err := http.NewRequest("GET", "?shortlink=missing", nil)
	if err != nil {
		t.Errorf("error creating request")
	}

	rr := httptest.NewRecorder()
	handler.redirectHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("wrong status code")
	}
}

func Test_viperlookuphandlergoodurlgoodkey(t *testing.T) {
	initviperconfig(t)

	handler := NewDictLookupHandler("")
	if handler == nil {
		t.Errorf("handler is nil")
	}

	req, err := http.NewRequest("GET", "/key1", nil)
	if err != nil {
		t.Errorf("error creating request")
	}

	rr := httptest.NewRecorder()
	handler.redirectHandler(rr, req)

	if rr.Code != http.StatusMovedPermanently {
		t.Errorf("wrong status code")
	}

	h := rr.Header().Get("location")

	if h != "/value1" {
		t.Errorf("wrong location")
	}
}

func Test_testReady(t *testing.T) {
	viper.Reset()
	initviperconfig(t)

	handler := NewDictLookupHandler("")
	if handler == nil {
		t.Errorf("handler is nil")
	}

	req, err := http.NewRequest("GET", "/ready", nil)
	if err != nil {
		t.Errorf("error creating request")
	}

	rr := httptest.NewRecorder()
	handler.readyHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("wrong status code expected %v, received %v", http.StatusOK, rr.Code)
	}
}

func Test_testNotReady(t *testing.T) {
	viper.Reset()
	initviperbadconfig(t)

	handler := NewDictLookupHandler("")
	if handler == nil {
		t.Errorf("handler is nil")
	}

	req, err := http.NewRequest("GET", "/ready", nil)
	if err != nil {
		t.Errorf("error creating request")
	}

	rr := httptest.NewRecorder()
	handler.readyHandler(rr, req)

	if rr.Code != http.StatusServiceUnavailable {
		t.Errorf("wrong status code expected %v, received %v", http.StatusServiceUnavailable, rr.Code)
	}
}

func Test_liveendtoend(t *testing.T) {
	initviperconfig(t)

	mux := http.NewServeMux()

	h := NewDictLookupHandler("")
	h.RegisterHandlers(mux)
	var wrappedmux http.Handler
	wrappedmux = basicserver.NewLoggingMiddleware(mux)
	wrappedmux = basicserver.NewPromMetricsMiddlewareWeb("pocketshorten_frontend", wrappedmux)

	s := httptest.NewServer(wrappedmux)
	defer s.Close()

	resp, err := http.DefaultClient.Get(s.URL + "/ready")
	if err != nil {
		t.Fatalf("Error")
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected response: %v", resp.StatusCode)
	}
}

func Test_shortenhandler(t *testing.T) {
	initviperconfig(t)

	mux := http.NewServeMux()

	h := NewDictLookupHandler("")
	h.RegisterHandlers(mux)
	var wrappedmux http.Handler
	wrappedmux = basicserver.NewLoggingMiddleware(mux)
	wrappedmux = basicserver.NewPromMetricsMiddlewareWeb("pocketshorten_frontend", wrappedmux)

	s := httptest.NewServer(wrappedmux)
	defer s.Close()

	resp, err := http.DefaultClient.Get(s.URL + "/hn")
	if err != nil {
		t.Fatalf("Error")
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected response: %v", resp.StatusCode)
	}

	//	if sr.Status() != http.StatusMovedPermanently {
	//		t.Fatalf("Statusmiddleware didn't work.  Expected 200 received %v", sr.Status())
	//	}
}

func initviperconfig(t *testing.T) {
	cpath := "testfiles/.pocketshorten.json"
	if internal.GitRoot == "" {
		t.Fatalf("GitRoot is empty, did you call SetupGitRoot() in test?")
	}
	configpath := path.Join(internal.GitRoot, cpath)
	internal.InitConfig(&configpath)
}

func initviperbadconfig(t *testing.T) {
	cpath := "testfiles/.pocketshorten_corrupt.json"
	if internal.GitRoot == "" {
		t.Fatalf("GitRoot is empty, did you call SetupGitRoot() in test?")
	}
	configpath := path.Join(internal.GitRoot, cpath)
	internal.InitConfig(&configpath)
}
