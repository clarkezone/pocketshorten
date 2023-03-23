/*
Copyright © 2022 clarkezone
*/
package main

import (
	clarkezoneLog "github.com/clarkezone/boosted-go/log"
	"github.com/sirupsen/logrus"

	"github.com/clarkezone/pocketshorten/cmd"
)

func main() {
	//clarkezoneLog.Init(logrus.WarnLevel)
	clarkezoneLog.Init(logrus.DebugLevel)
	cmd.Execute()
}
