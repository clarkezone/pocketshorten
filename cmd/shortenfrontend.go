// Package cmd contains the cli command definitions for pocketshorten
package cmd

/*
Copyright © 2022 James Clarke james@clarkezone.net

*/

import (
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/clarkezone/boosted-go/basicserverhttp"
	"github.com/clarkezone/boosted-go/middlewarehttp"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"

	"github.com/clarkezone/pocketshorten/internal"
	"github.com/clarkezone/pocketshorten/pkg/config"
	"github.com/clarkezone/pocketshorten/pkg/shortener"
)

const (
	prefix string = "pocketshorten_frontend"
)

// ShortenFrontendCmdState object
type ShortenFrontendCmdState struct {
	webserver *basicserverhttp.BasicServer
	shortener *shortener.ShortenHandler
}

func newShortenFrontend(parent *cobra.Command) (*ShortenFrontendCmdState, error) {
	ss := basicserverhttp.CreateBasicServer()
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

			mux := basicserverhttp.DefaultMux()
			cmdstate.shortener = shortener.NewDictLookupHandler(prefix)
			cmdstate.shortener.RegisterHandlers(mux)

			//sg := basicserver.NewStatusRecorder()
			var wrappedmux http.Handler
			wrappedmux = middlewarehttp.NewLoggingMiddleware(mux)
			wrappedmux = middlewarehttp.NewPromMetricsMiddlewareWeb(prefix, wrappedmux)
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

func (state *ShortenFrontendCmdState) configFlags(cmd *cobra.Command) error {
	cmd.PersistentFlags().StringVarP(&internal.ServiceURL, internal.ServiceURLVar, "",
		viper.GetString(internal.ServiceURLVar), "If value passed, testserverweb will delegate to this service")
	err := viper.BindPFlag(internal.ServiceURLVar, cmd.PersistentFlags().Lookup(internal.ServiceURLVar))
	if err != nil {
		return err
	}
	return nil
}
