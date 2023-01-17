package shortener

import (
	"context"

	"github.com/clarkezone/pocketshorten/internal"
	"github.com/clarkezone/pocketshorten/pkg/cacheloaderservice"
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpcStore
// TODO rename to dictCachePopulator
//
//lint:ignore U1000 reason backend not selected
type grpcStore struct {
	serviceURL string
	conn       *grpc.ClientConn
}

func (store *grpcStore) Store(short string, long string) error {
	return nil
}
func (store *grpcStore) Lookup(short string) (string, error) {
	return "", nil
}
func (store *grpcStore) Connect() error {
	return nil
}

func (store *grpcStore) startGrpcPopulate(errch chan error) {
	//TODO rename proto etc for uniform naming
	client := cacheloaderservice.NewUrlShortlinkCacheClient(store.conn)
	// this will block
	getitemsclient, err := client.GetItems(context.Background(), &cacheloaderservice.Empty{})
	if err != nil {
		clarkezoneLog.Errorf("grpcStore startGrpcPopulate error %v", err)
		errch <- err
	}
	// TODO while there are more items
	n, err := getitemsclient.Recv()
	if err != nil {
		clarkezoneLog.Errorf("grpcStore startGrpcPopulate error %v", err)
		errch <- err
		//TODO send error to channel
		//TODO handle reconnect
	}
	clarkezoneLog.Debugf("grpcStore startGrpcPopulate got %v", n)
	// TODO add to cache in thread safe manner
	clarkezoneLog.Debugf("grpcStore goroutine exited")
	// TODO kill goroutine on defer
	// TODO unit tests for dictstore and cachepopulator
	close(errch)
}

// TODO takes a dictstore, doesn't implement urlLookupService
func NewGrpcStore(u string) (*grpcStore, error) {
	ds := &grpcStore{}
	ds.serviceURL = u
	var err error
	ds.conn, err = grpc.Dial(internal.ServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer ds.conn.Close()
	errch := make(chan error)
	go ds.startGrpcPopulate(errch)
	//TODO how do we process errors on the errorchan
	//err = <-errch
	return ds, nil
}
