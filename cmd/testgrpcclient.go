// Package cmd contains the cli command definitions for previewd:w
package cmd

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

import (
	"context"

	"github.com/clarkezone/pocketshorten/internal"
	"github.com/clarkezone/pocketshorten/pkg/config"
	"github.com/clarkezone/pocketshorten/pkg/greetingservice"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

			clarkezoneLog.Successf("ServiceURL %v", internal.ServiceURL)

			conn, err := grpc.Dial(internal.ServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				clarkezoneLog.Errorf("fail to dial: %v", err)
			}
			defer conn.Close()

			if err == nil {
				client := greetingservice.NewGreeterClient(conn)
				result, err := client.GetGreeting(context.Background(), &greetingservice.Empty{})
				if err != nil {
					clarkezoneLog.Errorf("Error %v", err)
				} else {
					clarkezoneLog.Successf("Result %v", result.Name+result.Greeting)
				}
			}

			return err
		},
	}
	partent.AddCommand(cmd)
	cmd.PersistentFlags().StringVarP(&internal.ServiceURL, internal.ServiceURLVar, "",
		viper.GetString(internal.ServiceURLVar), "If value passed, testserverweb will delegate to this service")
	err := viper.BindPFlag(internal.ServiceURLVar, cmd.PersistentFlags().Lookup(internal.ServiceURLVar))
	if err != nil {
		panic(err)
	}
	return nil, nil
}
