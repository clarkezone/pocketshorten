// Package cmd contains the cli command definitions for pocketshorten
package cmd

/*
Copyright © 2022 James Clarke james@clarkezone.net

*/

import (
	"net/http"

	"github.com/clarkezone/pocketshorten/internal"
	"github.com/clarkezone/pocketshorten/pkg/basicserver"
	"github.com/clarkezone/pocketshorten/pkg/config"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/clarkezone/pocketshorten/pkg/shortener"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ShortenFrontendCmdState object
type ShortenFrontendCmdState struct {
	webserver *basicserver.BasicServer
	shortener *shortener.ShortenHandler
}

func newShortenFrontend(parent *cobra.Command) (*ShortenFrontendCmdState, error) {
	ss := basicserver.CreateBasicServer()
	cmdstate := &ShortenFrontendCmdState{webserver: ss, shortener: nil}

	// shortenservercmd represents the testserver command
	shortenservercmd := &cobra.Command{
		Use:   "servefrontend",
		Short: "Starts a pocketshorten server frontend instance",
		Long: `Starts a URL shorten server that will redirection
based on the url passed in:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clarkezoneLog.Successf("pocketshorten frontend server version %v,%v \n",
				config.VersionString, config.VersionHash)
			clarkezoneLog.Successf("Log level set to %v", internal.LogLevel)

			var err error
			cmdstate.shortener, err = shortener.NewGrpcLookupHandler(internal.ServiceURL)
			if err != nil {
				return err
			}

			mux := basicserver.DefaultMux()
			cmdstate.shortener.RegisterHandlers(mux)

			var wrappedmux http.Handler
			wrappedmux = basicserver.NewLoggingMiddleware(mux)
			wrappedmux = basicserver.NewPromMetricsMiddlewareWeb("pocketshorten_frontend", wrappedmux)

			if viper.GetString(internal.ServiceURLVar) != "" {
				clarkezoneLog.Successf("Delegating to %v", internal.ServiceURL)
			} else {
				clarkezoneLog.Debugf("Not delegating to %v", internal.ServiceURL)
			}

			clarkezoneLog.Successf("Starting pocketshorten frontend server on port %v", internal.Port)
			ss.StartMetrics()
			clarkezoneLog.Successf("Starting metrics on port %v", internal.MetricsPort)
			ss.StartListen("", wrappedmux)
			return ss.WaitforInterupt()
		},
	}
	err := cmdstate.configFlags(shortenservercmd)
	if err != nil {
		return nil, err
	}
	parent.AddCommand(shortenservercmd)
	return cmdstate, nil
}

func (ff *ShortenFrontendCmdState) configFlags(cmd *cobra.Command) error {
	cmd.PersistentFlags().StringVarP(&internal.ServiceURL, internal.ServiceURLVar, "",
		viper.GetString(internal.ServiceURLVar), "If value passed, testserverweb will delegate to this service")
	err := viper.BindPFlag(internal.ServiceURLVar, cmd.PersistentFlags().Lookup(internal.ServiceURLVar))
	if err != nil {
		return err
	}
	return nil
}
