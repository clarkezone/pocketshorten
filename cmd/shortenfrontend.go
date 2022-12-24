// Package cmd contains the cli command definitions for pocketshorten
package cmd

/*
Copyright © 2022 James Clarke james@clarkezone.net

*/

import (
	"fmt"
	"net/http"

	"github.com/clarkezone/pocketshorten/internal"
	"github.com/clarkezone/pocketshorten/pkg/basicserver"
	"github.com/clarkezone/pocketshorten/pkg/config"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var shortenserver = basicserver.CreateBasicServer()

var (
	// shortenservercmd represents the testserver command
	shortenservercmd = &cobra.Command{
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
			mux := basicserver.DefaultMux()
			mux.HandleFunc("/", getRedirectHandler())

			var wrappedmux http.Handler
			wrappedmux = basicserver.NewLoggingMiddleware(mux)
			wrappedmux = basicserver.NewPromMetricsMiddlewareWeb("pocketshorten_frontend", wrappedmux)

			if viper.GetString(internal.ServiceURLVar) != "" {
				clarkezoneLog.Successf("Delegating to %v", internal.ServiceURL)
			} else {
				clarkezoneLog.Debugf("Not delegating to %v", internal.ServiceURL)
			}

			clarkezoneLog.Successf("Starting pocketshorten frontend server on port %v", internal.Port)
			shortenserver.StartMetrics()
			clarkezoneLog.Successf("Starting metrics on port %v", internal.MetricsPort)
			shortenserver.StartListen("", wrappedmux)
			return shortenserver.WaitforInterupt()
		},
	}
)

func getRedirectHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		requested := r.URL.Query().Get("shortlink")

		if requested == "" {
			writeOutput(w, "please supply shortlink query parameter")
			return
		}

		if requested == "james" {
			http.Redirect(w, r, "https://github.com/clarkezone", http.StatusMovedPermanently)
			return
		}

		writeOutput(w, fmt.Sprintf("shortlink %v notfound", requested))
	}
}

func writeOutput(w http.ResponseWriter, message string) {
	_, err := w.Write([]byte(message))
	if err != nil {
		clarkezoneLog.Debugf("Failed to write bytes %v\n", err)
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(shortenservercmd)
	shortenservercmd.PersistentFlags().StringVarP(&internal.ServiceURL, internal.ServiceURLVar, "",
		viper.GetString(internal.ServiceURLVar), "If value passed, testserverweb will delegate to this service")
	err := viper.BindPFlag(internal.ServiceURLVar, shortenservercmd.PersistentFlags().Lookup(internal.ServiceURLVar))
	if err != nil {
		panic(err)
	}
}
