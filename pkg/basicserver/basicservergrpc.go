// Package basicserver contains a dummy server implementation for testing metrics and logging
package basicserver

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/clarkezone/pocketshorten/internal"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

// type cleanupfunc func()

// BasicServerGrpc object
type BasicServerGrpc struct {
	lis             *net.Listener
	grpcServer      *grpc.Server
	ctx             context.Context
	cancel          context.CancelFunc
	exitchan        chan (bool)
	metricsserver   *http.Server
	metricsexitchan chan (bool)
}

// CreateBasicServerGrpc Create BasicServer object and return
func CreateBasicServerGrpc() *BasicServerGrpc {
	bs := BasicServerGrpc{}
	return &bs
}

// StartListen Start listening for a connection
func (bs *BasicServerGrpc) StartListen(secret string) *grpc.Server {
	clarkezoneLog.Successf("starting... basic server on :%v", fmt.Sprint(internal.Port))

	bs.exitchan = make(chan bool)
	bs.ctx, bs.cancel = context.WithCancel(context.Background())
	lis, err := net.Listen("tcp", ":"+fmt.Sprint(internal.Port))
	if err != nil {
		panic(err)
	}

	bs.lis = &lis
	var opts []grpc.ServerOption
	bs.grpcServer = grpc.NewServer(opts...)

	go func() {
		err = bs.grpcServer.Serve(*bs.lis)
		if err.Error() != "http: Server closed" {
			panic(err)
		}
		defer func() {
			clarkezoneLog.Debugf("Webserver goroutine exited")
			bs.exitchan <- true
		}()
	}()
	return bs.grpcServer
}

// StartMetrics Start listening for a connection for metrics
func (bs *BasicServerGrpc) StartMetrics() {
	clarkezoneLog.Successf("starting... metrics on :%v", fmt.Sprint(internal.MetricsPort))

	if bs.ctx == nil {
		bs.ctx, bs.cancel = context.WithCancel(context.Background())
	}

	bs.metricsexitchan = make(chan bool)
	bs.metricsserver = &http.Server{Addr: ":" + fmt.Sprint(internal.MetricsPort)}
	bs.metricsserver.Handler = promhttp.Handler()

	go func() {
		err := bs.metricsserver.ListenAndServe()
		if err.Error() != "http: Server closed" {
			panic(err)
		}
		defer func() {
			clarkezoneLog.Debugf("Metrics server goroutine exited")
			bs.metricsexitchan <- true
		}()
	}()
}

// WaitforInterupt Wait for a sigterm event or for user to press control c when running interacticely
func (bs *BasicServerGrpc) WaitforInterupt() error {
	if bs.exitchan == nil {
		clarkezoneLog.Debugf("WaitForInterupt(): server not started\n")
		return fmt.Errorf("server not started")
	}
	ch := make(chan struct{})
	handleSig(func() { close(ch) })
	clarkezoneLog.Successf("Waiting for user to press control c or sig terminate\n")
	<-ch
	clarkezoneLog.Debugf("Terminate signal detected, closing job manager\n")
	return bs.Shutdown()
}

// func handleSig(cleanupwork cleanupfunc) chan struct{} {
// 	signalChan := make(chan os.Signal, 1)
// 	cleanupDone := make(chan struct{})
// 	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
// 	go func() {
// 		<-signalChan
// 		clarkezoneLog.Debugf("\nhandleSig Received an interrupt, stopping services...\n")
// 		if cleanupwork != nil {
// 			clarkezoneLog.Debugf("cleanup work found")
// 			cleanupwork()
// 			clarkezoneLog.Debugf("cleanup work completed")
// 		}

// 		close(cleanupDone)
// 	}()
// 	return cleanupDone
// }

// Shutdown terminates the listening thread
func (bs *BasicServerGrpc) Shutdown() error {
	if bs.exitchan == nil {
		clarkezoneLog.Debugf("BasicServer: no exit channel detected on shutdown\n")
		return fmt.Errorf("BasicServer: no exit channel detected on shutdown")
	}
	defer bs.ctx.Done()
	defer bs.cancel()
	clarkezoneLog.Debugf("BasicServer: request httpserver shutdown")
	bs.grpcServer.GracefulStop()
	clarkezoneLog.Debugf("BasicServer: shutdwon completed, wait for exitchan")
	<-bs.exitchan
	clarkezoneLog.Debugf("BasicServer: exit completed function returqn")

	if bs.metricsserver != nil {
		if bs.metricsexitchan == nil {
			clarkezoneLog.Debugf("BasicServer: metrics no exit channel detected on shutdown\n")
			return fmt.Errorf("BasicServer: metrics no exit channel detected on shutdown")
		}
		metricsexit := bs.metricsserver.Shutdown(bs.ctx)

		clarkezoneLog.Debugf("BasicServer: metrics shutdwon completed, wait for exitchan")
		<-bs.metricsexitchan
		clarkezoneLog.Debugf("BasicServer: metrics exit completed function returqn")
		if metricsexit != nil {
			return metricsexit
		}
	} else {

		clarkezoneLog.Debugf("BasicServer: no metrics server detected on shutdown hence skipping extichannel\n")
	}
	return nil
}
