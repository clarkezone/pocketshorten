// Package greetingservice is an implementation of the GreetingService service.
package greetingservice

import context "context"

// GreetingServer is the server API for GreetingService service.
type GreetingServer struct {
	UnimplementedGreeterServer
}

// GetGreeting implements GreetingServer
func (s *GreetingServer) GetGreeting(ctx context.Context, in *Empty) (*Greeting, error) {
	// Do something with the input message
	// ...

	// Return a response message
	return &Greeting{}, nil
}
