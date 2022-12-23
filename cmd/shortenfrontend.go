// Package cmd contains the cli command definitions for pocketshorten
package cmd

/*
Copyright © 2022 James Clarke james@clarkezone.net

*/

import (
	"context"
	"fmt"
	"net/http"

	"github.com/clarkezone/pocketshorten/internal"
	"github.com/clarkezone/pocketshorten/pkg/basicserver"
	"github.com/clarkezone/pocketshorten/pkg/config"
	"github.com/clarkezone/pocketshorten/pkg/greetingservice"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var shortenserver = basicserver.CreateBasicServer()

var (
	// shortenservercmd represents the testserver command
	shortenservercmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts a test pocketshorten server to test logging and metrics",
		Long: `Starts a URL shorten server that will redirection
based on the url passed in:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clarkezoneLog.Successf("pocketshorten server  version %v,%v \n",
				config.VersionString, config.VersionHash)
			clarkezoneLog.Successf("Log level set to %v", internal.LogLevel)
			mux := basicserver.DefaultMux()
			mux.HandleFunc("/", getServerHandler())

			var wrappedmux http.Handler
			wrappedmux = basicserver.NewLoggingMiddleware(mux)
			wrappedmux = basicserver.NewPromMetricsMiddlewareWeb("pocketshortener_testWebservice", wrappedmux)

			if viper.GetString(internal.ServiceURLVar) != "" {
				clarkezoneLog.Successf("Delegating to %v", internal.ServiceURL)
			} else {
				clarkezoneLog.Debugf("Not delegating to %v", internal.ServiceURL)
			}

			clarkezoneLog.Successf("Starting pocketshorten server on port %v", internal.Port)
			shortenserver.StartMetrics()
			clarkezoneLog.Successf("Starting metrics on port %v", internal.MetricsPort)
			shortenserver.StartListen("", wrappedmux)
			return shortenserver.WaitforInterupt()
		},
	}
)

func getServerHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var message string

		if viper.GetString(internal.ServiceURLVar) != "" {
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
					clarkezoneLog.Successf("Result %v from %v", result.Greeting, result.Name)
					message = fmt.Sprintf("%v from %v at %v<br>", result.Name, result.Greeting, result.LastUpdated)
				}
			} else {
				clarkezoneLog.Errorf("Error %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				_, err := w.Write([]byte("500 - Something bad happened!"))
				if err != nil {
					clarkezoneLog.Errorf("Error %v", err)
				}
			}
		} else {
			message = fmt.Sprintln("Hello World<br>")
		}

		_, err := w.Write([]byte(message))
		if err != nil {
			clarkezoneLog.Debugf("Failed to write bytes %v\n", err)
			panic(err)
		}

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