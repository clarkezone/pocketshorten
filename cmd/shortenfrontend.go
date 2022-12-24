// Package cmd contains the cli command definitions for pocketshorten
package cmd

/*
Copyright © 2022 James Clarke james@clarkezone.net

*/

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/clarkezone/pocketshorten/internal"
	"github.com/clarkezone/pocketshorten/pkg/basicserver"
	"github.com/clarkezone/pocketshorten/pkg/config"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ShortenFrontendCmd object
type ShortenFrontendCmd struct {
	bs *basicserver.BasicServer
	lh *lookupHandler
}

func newShortenFrontend(parent *cobra.Command) (*ShortenFrontendCmd, error) {

	lhandler := newLookupHandler()
	ss := basicserver.CreateBasicServer()
	sfc := &ShortenFrontendCmd{bs: ss, lh: lhandler}
	// TODO populate dictstore
	// TODO populate dictstore from grpc

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

			mux := basicserver.DefaultMux()
			mux.HandleFunc("/", sfc.lh.redirectHandler)

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
	err := sfc.configFlags(shortenservercmd)
	if err != nil {
		return nil, err
	}
	parent.AddCommand(shortenservercmd)
	return sfc, nil
}

func (ff *ShortenFrontendCmd) configFlags(cmd *cobra.Command) error {
	cmd.PersistentFlags().StringVarP(&internal.ServiceURL, internal.ServiceURLVar, "",
		viper.GetString(internal.ServiceURLVar), "If value passed, testserverweb will delegate to this service")
	err := viper.BindPFlag(internal.ServiceURLVar, cmd.PersistentFlags().Lookup(internal.ServiceURLVar))
	if err != nil {
		return err
	}
	return nil
}

type urlLookupService interface {
	Store(string, string) error
	Lookup(string) (string, error)
}

type dictStore struct {
	m map[string]string
}

func (store *dictStore) Store(short string, long string) error {
	//TODO logging, telemetry
	store.m[short] = long
	return nil
}

func (store *dictStore) Lookup(short string) (string, error) {
	//TODO logging, telemetry
	val, pr := store.m[short]
	if pr {
		return val, nil
	}
	return "", errors.New("Key not found")
}

func newLookupHandler() *lookupHandler {
	ds := &dictStore{}
	ds.m = make(map[string]string)
	lh := &lookupHandler{storage: ds}
	return lh
}

type lookupHandler struct {
	storage urlLookupService
}

func (lh *lookupHandler) redirectHandler(w http.ResponseWriter, r *http.Request) {
	//TODO logging, telemetry
	requested := r.URL.Query().Get("shortlink")

	if requested == "" {
		writeOutputError(w, "please supply shortlink query parameter")
		return
	}
	uri, err := lh.storage.Lookup(requested)
	if err != nil {
		writeOutputError(w, fmt.Sprintf("shortlink %v notfound", requested))
		return
	} else {
		clarkezoneLog.Debugf("redirecting to %v", uri)
	}

	http.Redirect(w, r, uri, http.StatusMovedPermanently)

}

func writeOutputError(w http.ResponseWriter, message string) {
	clarkezoneLog.Debugf("Error reported to user %v", message)
	_, err := w.Write([]byte(message))
	if err != nil {
		clarkezoneLog.Debugf("writeOutputError: Failed to write bytes %v\n", err)
	}
	w.WriteHeader(http.StatusNotFound)
}
