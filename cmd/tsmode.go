package cmd

import "fmt"

type modeValue string

const (
	httpserver modeValue = "httpserver"
	grpcserver modeValue = "grpcserver"
	grpcclient modeValue = "grpcclient"
)

func (m *modeValue) String() string {
	return string(*m)
}

func (m *modeValue) Set(value string) error {
	if value != string(httpserver) && value != string(grpcserver) && value != string(grpcclient) {
		return fmt.Errorf("invalid value for mode: %s", value)
	}
	*m = modeValue(value)
	return nil
}

func (m *modeValue) Type() string {
	return "mode"
}
