package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/clarkezone/pocketshorten/internal"
	"github.com/clarkezone/pocketshorten/pkg/config"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

// TestMain initizlie all tests
func TestMain(m *testing.M) {
	internal.SetupGitRoot()
	clarkezoneLog.Init(logrus.DebugLevel)
	code := m.Run()
	os.Exit(code)
}

func Test_ExecuteVersion(t *testing.T) {
	config.VersionString = "1"
	config.VersionHash = "A"
	cmd := getVersionCommand()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expected := "pocketshorten version:1 hash:A\n"
	if string(out) != expected {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, string(out))
	}
}
