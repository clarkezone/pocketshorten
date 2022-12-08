/*
Copyright © 2022 clarkezone
*/
package main

import (
	"github.com/clarkezone/pocketshorten/cmd"
	clarkezoneGreeter "github.com/clarkezone/pocketshorten/pkg/greetingservice"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/sirupsen/logrus"
)

func main() {
	clarkezoneGreeter.NewGreeterClient(nil)
	clarkezoneLog.Init(logrus.WarnLevel)
	cmd.Execute()
}
