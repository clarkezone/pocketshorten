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

// Grpc object
type Grpc struct {
	lis             *net.Listener
	grpcServer      *grpc.Server
	ctx             context.Context
	cancel          context.CancelFunc
	exitchan        chan (bool)
	metricsserver   *http.Server
	metricsexitchan chan (bool)
	interceptors    []grpc.UnaryServerInterceptor
}

// CreateGrpc Create BasicServer object and return
func CreateGrpc() *Grpc {
	bs := Grpc{}
	return &bs
}

// AddUnaryInterceptor adds middleware to chain
func (bs *Grpc) AddUnaryInterceptor(i grpc.UnaryServerInterceptor) {
	bs.interceptors = append(bs.interceptors, i)
}

// StartListen Start listening for a connection
func (bs *Grpc) StartListen(secret string) *grpc.Server {
	clarkezoneLog.Successf("starting... basic grpc server on :%v", fmt.Sprint(internal.Port))

	bs.exitchan = make(chan bool)
	bs.ctx, bs.cancel = context.WithCancel(context.Background())
	lis, err := net.Listen("tcp", "0.0.0.0:"+fmt.Sprint(internal.Port))
	if err != nil {
		panic(err)
	}

	bs.lis = &lis

	bs.interceptors = append(bs.interceptors, bs.logsUnaryInterceptor)

	opts := []grpc.ServerOption{grpc.ChainUnaryInterceptor(
		bs.interceptors...)}

	clarkezoneLog.Debugf("new grpc server with %v interceptors", len(bs.interceptors))

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
func (bs *Grpc) StartMetrics() {
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
func (bs *Grpc) WaitforInterupt() error {
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
func (bs *Grpc) Shutdown() error {
	if bs.exitchan == nil {
		clarkezoneLog.Debugf("Grpc: no exit channel detected on shutdown\n")
		return fmt.Errorf("Grpc: no exit channel detected on shutdown")
	}
	defer bs.ctx.Done()
	defer bs.cancel()
	clarkezoneLog.Debugf("Grpc: request httpserver shutdown")
	bs.grpcServer.GracefulStop()
	clarkezoneLog.Debugf("Grpc: shutdwon completed, wait for exitchan")
	<-bs.exitchan
	clarkezoneLog.Debugf("Grpc: exit completed function returqn")

	if bs.metricsserver != nil {
		if bs.metricsexitchan == nil {
			clarkezoneLog.Debugf("Grpc: metrics no exit channel detected on shutdown\n")
			return fmt.Errorf("Grpc: metrics no exit channel detected on shutdown")
		}
		metricsexit := bs.metricsserver.Shutdown(bs.ctx)

		clarkezoneLog.Debugf("Grpc: metrics shutdwon completed, wait for exitchan")
		<-bs.metricsexitchan
		clarkezoneLog.Debugf("Grpc: metrics exit completed function returqn")
		if metricsexit != nil {
			return metricsexit
		}
	} else {

		clarkezoneLog.Debugf("Grpc: no metrics server detected on shutdown hence skipping extichannel\n")
	}
	return nil
}

func (bs *Grpc) logsUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	clarkezoneLog.Debugf("gRPC method called %v", info.FullMethod)
	return handler(ctx, req)
}
