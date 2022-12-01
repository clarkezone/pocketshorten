// Package cmd contains the cli command definitions for previewd:w
package cmd

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

import (
	"fmt"
	"net/http"

	"github.com/clarkezone/pocketshorten/pkg/basicserver"
	"github.com/clarkezone/pocketshorten/pkg/config"
	"github.com/spf13/cobra"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

var bs = basicserver.CreateBasicServer()

var (
	// testserverCmd represents the testserver command
	testserverCmd = &cobra.Command{
		Use:   "testserver",
		Short: "Starts a test server to test logging and metrics",
		Long: `Starts a listener that will
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clarkezoneLog.Successf("previewd version %v,%v started in testserver mode\n",
				config.VersionString, config.VersionHash)
			mux := basicserver.DefaultMux()
			mux.HandleFunc("/", getHelloHandler())
			// mux.Handle("/metrics", promhttp.Handler())

			var wrappedmux http.Handler
			wrappedmux = basicserver.NewLoggingMiddleware(mux)
			wrappedmux = basicserver.NewPromMetricsMiddleware("previewd_testserver_", wrappedmux)

			bs.StartListen("", wrappedmux)
			return bs.WaitforInterupt()
		},
	}
)

func getHelloHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintln("Hello World<BR>")
		_, err := w.Write([]byte(message))
		if err != nil {
			clarkezoneLog.Debugf("Failed to write bytes %v\n", err)
			panic(err)
		}
	}
}

func init() {
	rootCmd.AddCommand(testserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
