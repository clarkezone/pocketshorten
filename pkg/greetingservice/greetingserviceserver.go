// Package greetingservice is an implementation of the GreetingService service.
package greetingservice

import (
	context "context"
	"os"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// GreetingServer is the server API for GreetingService service.
type GreetingServer struct {
	UnimplementedGreeterServer
}

// GetGreeting implements GreetingServer
func (s *GreetingServer) GetGreeting(ctx context.Context, in *Empty) (*Greeting, error) {
	name := os.Getenv("MY_POD_NAME")
	return &Greeting{
		Name:        name,
		Greeting:    "Hello",
		LastUpdated: timestamppb.Now(),
	}, nil
}
