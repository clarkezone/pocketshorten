/*
Copyright © 2022 clarkezone
*/
package main

import (
	"github.com/clarkezone/pocketshorten/cmd"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/sirupsen/logrus"
)

func main() {
	clarkezoneLog.Init(logrus.WarnLevel)
	cmd.Execute()
}
