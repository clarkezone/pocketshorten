package shortener

import (
	"os"
	"testing"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/sirupsen/logrus"
)

// TestMain initizlie all tests
func TestMain(m *testing.M) {
	clarkezoneLog.Init(logrus.DebugLevel)
	code := m.Run()
	os.Exit(code)
}

func Test_webhooklistening(t *testing.T) {
	handler := newDictLookupHandler()
	if handler == nil {
		t.Errorf("handler is nil")
	}
}
