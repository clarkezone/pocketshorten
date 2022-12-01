package internal

import "testing"

func Test_Default(t *testing.T) {
	if Port != 8090 {
		t.Errorf("default wrong")
	}
}
