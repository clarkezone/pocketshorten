package shortener

import (
	"os"
	"path"
	"testing"

	"github.com/clarkezone/pocketshorten/internal"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/sirupsen/logrus"
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
	handler := NewDictStore(tlh)
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
	handler := newDictLookupHandler()
	if handler == nil {
		t.Errorf("handler is nil")
	}
}

func Test_viperlookuphandler(t *testing.T) {
	initviperconfig(t)

	handler := newDictLookupHandler()
	if handler == nil {
		t.Errorf("handler is nil")
	}

	//lint:ignore SA5011 reason test
	if handler.storage.Count() != 3 {
		t.Errorf("wrong number of items in storage")
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
