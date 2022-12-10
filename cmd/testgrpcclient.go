// Package cmd contains the cli command definitions for previewd:w
package cmd

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

import (
	"github.com/clarkezone/pocketshorten/internal"
	"github.com/clarkezone/pocketshorten/pkg/config"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/spf13/cobra"
)

// TestClientGrpcCmd is the command to start a test grpc client
type TestClientGrpcCmd struct {
}

func newTestClientGrpcCmd(partent *cobra.Command) (*TestClientGrpcCmd, error) {
	cmd := &cobra.Command{
		Use:   "testclientgrpc",
		Short: "Starts a client",
		Long: `Starts a client that will call the testservergrpc which must be running
have already been started with the testservergrpc command`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clarkezoneLog.Successf("pocketshorten version %v,%v started in testclientgrpc mode\n",
				config.VersionString, config.VersionHash)
			clarkezoneLog.Successf("Log level set to %v", internal.LogLevel)
			return nil
		},
	}
	partent.AddCommand(cmd)
	return nil, nil
}
