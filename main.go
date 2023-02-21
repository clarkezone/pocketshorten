/*
Copyright © 2022 clarkezone
*/
package main

import (
	"github.com/clarkezone/boosted-go/testmod"
	"github.com/sirupsen/logrus"

	"github.com/clarkezone/pocketshorten/cmd"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

func main() {
	testmod.CreateThing()
	//clarkezoneLog.Init(logrus.WarnLevel)
	clarkezoneLog.Init(logrus.DebugLevel)
	cmd.Execute()
}
