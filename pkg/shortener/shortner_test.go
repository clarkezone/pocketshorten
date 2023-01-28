package shortener

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/clarkezone/pocketshorten/internal"
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
	err := ls.Store("one", "two")
	err2 := ls.Store("three", "four")
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
		if handler.storage.Count() != 3 {
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

func Test_testReady(t *testing.T) {
	viper.Reset()
	initviperconfig(t)

	handler := NewDictLookupHandler()
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

	handler := NewDictLookupHandler()
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
