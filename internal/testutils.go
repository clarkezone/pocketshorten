package internal

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
)

const (
	testreponame         = "TEST_GITLAYER_REPO_NOAUTHURL"
	testlocaldirname     = "TEST_GITLAYER_LOCALDIR"
	testbranchswitchname = "TEST_GITLAYER_BRANCHSWITCH"
	testsecurereponame   = "TEST_GITLAYER_SECURE_REPO_NOAUTH"
	//nolint
	testsecureclonepwname = "TEST_GITLAYER_SECURECLONEPWNAME"
)

// configure environment variables by:
// 1. command palette: open settings (json)
// 2. append the following
// "go.testEnvFile": "/home/james/.previewd_test.env",
// 3. contents of file
// TEST_GITLAYER_REPO_NOAUTHURL="https:/"
// TEST_GITLAYER_LOCALDIR=""
// TEST_GITLAYER_BRANCHSWITCH=""
// TEST_GITLAYER_SECURE_REPO_NOAUTH=""
// TEST_GITLAYER_SECURECLONEPW=""
// TEST_GITLAYER_TESTLOCALK8S=""

// Getenv returns environment variables for use in tests
func Getenv(t *testing.T) (string, string, string, string, string) {
	repo := os.Getenv(testreponame)
	localdr := os.Getenv(testlocaldirname)
	testbranchswitch := os.Getenv(testbranchswitchname)
	reposecure := os.Getenv(testsecurereponame)
	secureclonepw := os.Getenv(testsecureclonepwname)
	if repo == "" || localdr == "" || testbranchswitch == "" {
		t.Fatalf("Test environment variables not configured repo:%v, localdr:%v, testbranchswitch:%v,\n",
			repo, localdr, testbranchswitch)
	}
	return repo, localdr, testbranchswitch, reposecure, secureclonepw
}

// GitRoot is the root of current git repo used in testing
var GitRoot string

// SetupGitRoot finds the gitroot for this repo
func SetupGitRoot() {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")

	output, err := cmd.CombinedOutput()
	if err != nil {
		panic("couldn't read output from git command get gitroot")
	}
	GitRoot = string(output)
	GitRoot = strings.TrimSuffix(GitRoot, "\n")
}

// GetTestConfigPath returns a local testing config for k8s
func GetTestConfigPath(t *testing.T) string {
	if GitRoot == "" {
		t.Fatalf("GitRoot is empty, did you call SetupGitRoot() in test?")
	}
	configpath := path.Join(GitRoot, "integration/secrets/k3s-c2.yaml")
	return configpath
}
