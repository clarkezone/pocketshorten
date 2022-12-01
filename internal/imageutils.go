package internal

import (
	"fmt"
	"runtime"
)

// GetJekyllImage returns the path to the official render image
func GetJekyllImage() string {
	fmt.Printf("%v", runtime.GOARCH)
	return "registry.hub.docker.com/clarkezone/jekyll:sha-df0a146"
}

// GetJekyllCommands returns the command and params used to trigger a render job
func GetJekyllCommands() ([]string, []string) {
	command := []string{"sh", "-c", "--"}
	params := []string{"cd /src/source;bundle install;bundle exec jekyll build -d /site JEKYLL_ENV=production"}
	return command, params
}
