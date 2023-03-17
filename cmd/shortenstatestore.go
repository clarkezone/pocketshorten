// Package cmd contains the cli command definitions for gomicroservicestarter
package cmd

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

import (
	"github.com/clarkezone/boosted-go/basicservergrpc"
	"github.com/clarkezone/boosted-go/middlewaregrpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/clarkezone/pocketshorten/internal"
	"github.com/clarkezone/pocketshorten/pkg/cacheloaderservice"
	"github.com/clarkezone/pocketshorten/pkg/config"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

// ShortenStateStore is the command to start a test grpc server
type ShortenStateStore struct {
	bs  *basicservergrpc.Grpc
	mid *middlewaregrpc.PromMetricsMiddlewareGrpc
}

func newShortenStateStore(partent *cobra.Command) (*ShortenStateStore, error) {
	bssssGrpc := basicservergrpc.CreateGrpc()
	sss := &ShortenStateStore{
		bs: bssssGrpc,
	}
	cmd := &cobra.Command{
		Use:   "shortenstatestore",
		Short: "Starts a pocketshorten state store that stores urls / shortcuts mappings",
		Long: `A frontend pocketshorten instance needs a statestore as the source of truth for URLs:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clarkezoneLog.Successf("pocketshorten statestore %v,%v started\n",
				config.VersionString, config.VersionHash)
			clarkezoneLog.Successf("Log level set to %v", internal.LogLevel)

			values := viper.Get("values").([]interface{})

			if values == nil {
				clarkezoneLog.Debugf("Shortenstatestore Valus is nil: %v", values)
			} else {

				clarkezoneLog.Debugf("Shortenstatestore Valus is not nil: number in collection %v", len(values))
			}

			// Iterate over the string pairs in the array
			//	for _, pair := range values {
			//		key := pair[0]
			//		value := pair[1]
			//		clarkezoneLog.Debugf("%s: %s\n", key, value)
			//	}

			sss.mid = middlewaregrpc.NewPromMetricsMiddlewareGrpc("pocketshorten_statestore")
			bssssGrpc.AddUnaryInterceptor(sss.mid.MetricsUnaryInterceptor)
			clarkezoneLog.Successf("Starting pocketshorten statestore server on port %v", internal.Port)
			bssssGrpc.StartMetrics()
			clarkezoneLog.Successf("Starting metrics on port %v", internal.MetricsPort)
			serv := bssssGrpc.StartListen("")
			cls := &cacheloaderservice.CacheLoaderServer{}
			cacheloaderservice.RegisterUrlShortlinkCacheServer(serv, cls)
			return bssssGrpc.WaitforInterupt()
		},
	}
	err := sss.configFlags(cmd)
	if err != nil {
		return nil, err
	}
	partent.AddCommand(cmd)
	return sss, nil
}

func (ts *ShortenStateStore) configFlags(cmd *cobra.Command) error {
	m := modeValue(internal.StartupMode)

	cmd.PersistentFlags().VarP(&m, "startupmode", "", "startup mode (httpserver, grpcserver, grpcclient) (default is httpserver)")
	err := viper.BindPFlag("startupmode", cmd.PersistentFlags().Lookup(internal.StartupMode))

	if err != nil {
		return err
	}
	return nil
}
