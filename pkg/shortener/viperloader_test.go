package shortener

import (
	"path"
	"testing"

	"github.com/spf13/viper"

	"github.com/clarkezone/pocketshorten/internal"
)

func Test_LoadfromConfig(t *testing.T) {
	viper.Reset()
	cpath := "testfiles/.pocketshorten.json"
	if internal.GitRoot == "" {
		t.Fatalf("GitRoot is empty, did you call SetupGitRoot() in test?")
	}
	configpath := path.Join(internal.GitRoot, cpath)
	internal.InitConfig(&configpath)
	viperLoader := &viperLoader{}
	err := viperLoader.Init(nil)
	if err != nil {
		t.Fatalf("init failed with %v", err)
	}
}

func Test_LoadfromBadConfigFail(t *testing.T) {
	viper.Reset()
	cpath := "testfiles/.pocketshorten_corrupt.json"
	if internal.GitRoot == "" {
		t.Fatalf("GitRoot is empty, did you call SetupGitRoot() in test?")
	}
	configpath := path.Join(internal.GitRoot, cpath)
	internal.InitConfig(&configpath)
	vl := &viperLoader{}
	err := vl.Init(nil)
	if err == nil {
		t.Fatalf("init failed to detect corrupt config file with %v", err)
	}
}
