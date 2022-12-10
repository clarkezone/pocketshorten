// Package greetingservice is an implementation of the GreetingService service.
package greetingservice

import (
	context "context"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

// GreetingServer is the server API for GreetingService service.
type GreetingServer struct {
	UnimplementedGreeterServer
}

// GetGreeting implements GreetingServer
func (s *GreetingServer) GetGreeting(ctx context.Context, in *Empty) (*Greeting, error) {
	clarkezoneLog.Debugf("GetGreeting called")
	return &Greeting{
		Name:     "James",
		Greeting: "Hello",
	}, nil
}
