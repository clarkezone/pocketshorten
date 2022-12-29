// Package cacheloaderservice is an implementation of the GreetingService service.
package cacheloaderservice

// CacheLoaderServer is the server API for GreetingService service.
type CacheLoaderServer struct {
	UnimplementedUrlShortlinkCacheServer
	waiter chan bool
}

// GetItems implements CacheLoaderServer
func (s *CacheLoaderServer) GetItems(e *Empty, stream UrlShortlinkCache_GetItemsServer) (err error) {
	//TODO: should this be a goroute?
	for {
		u := UrlShortLink{}
		// Send the data to the client.
		if err := stream.Send(&u); err != nil {
			return err
		}
		<-s.waiter
	}
}
