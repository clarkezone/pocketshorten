package shortener

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/sirupsen/logrus"

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

func (tlh *testLookupHandler) Init(ls urlLookupService) {
	tlh.lookup = ls
	err := ls.Store("one", "two")
	err2 := ls.Store("three", "four")
	if err != nil || err2 != nil {
		panic("Init Failed")
	}
}

func Test_dictStore(t *testing.T) {
	// pass in viper loader
	tlh := &testLookupHandler{}
	handler := newDictStore(tlh)
	if handler == nil {
		t.Errorf("handler is nil")
	}
	val, err := handler.Lookup("three")
	if err != nil {
		t.Errorf("lookup failed")
	}
	if val != "four" {
		t.Errorf("Wrong value")
	}
}

func Test_dictLookupHandler(t *testing.T) {
	handler := NewDictLookupHandler()
	if handler == nil {
		t.Errorf("handler is nil")
	}
}

func Test_viperlookuphandlerinit(t *testing.T) {
	initviperconfig(t)

	handler := NewDictLookupHandler()
	if handler != nil {

		//lint:ignore SA5011 reason test
		if handler.storage.Count() != 5 {
			t.Errorf("wrong number of items in storage")
		}
	}
}

func Test_viperlookuphandlerlookupbadurl(t *testing.T) {
	initviperconfig(t)

	handler := NewDictLookupHandler()
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

func Test_viperlookuphandlergoodurlmissingkey(t *testing.T) {
	initviperconfig(t)

	handler := NewDictLookupHandler()
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

	handler := NewDictLookupHandler()
	if handler == nil {
		t.Errorf("handler is nil")
	}

	req, err := http.NewRequest("GET", "?shortlink=key1", nil)
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

func Test_shortenhandler(t *testing.T) {
	initviperconfig(t)

	mux := http.NewServeMux()

	h := NewDictLookupHandler()
	h.RegisterHandlers(mux)
	var wrappedmux http.Handler
	sr := basicserver.NewStatusRecorder()
	wrappedmux = basicserver.NewLoggingMiddleware(mux)
	wrappedmux = basicserver.NewPromMetricsMiddlewareWeb("pocketshorten_frontend", wrappedmux, sr)
	wrappedmux = basicserver.NewStatusMiddlewareWeb(wrappedmux, sr)

	s := httptest.NewServer(wrappedmux)
	defer s.Close()

	resp, err := http.DefaultClient.Get(s.URL + "?shortlink=hn")
	if err != nil {
		t.Fatalf("Error")
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected response: %v", resp.StatusCode)
	}

	if sr.Status() != http.StatusMovedPermanently {
		t.Fatalf("Statusmiddleware didn't work.  Expected 200 received %v", sr.Status())
	}

}

func initviperconfig(t *testing.T) {
	cpath := "testfiles/.pocketshorten.json"
	if internal.GitRoot == "" {
		t.Fatalf("GitRoot is empty, did you call SetupGitRoot() in test?")
	}
	configpath := path.Join(internal.GitRoot, cpath)
	internal.InitConfig(&configpath)
}
