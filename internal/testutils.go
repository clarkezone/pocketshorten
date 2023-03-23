package internal

import (
	"os/exec"
	"strings"
)

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
