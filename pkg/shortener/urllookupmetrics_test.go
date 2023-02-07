package shortener

import (
	"testing"

	"github.com/spf13/viper"
)

func Test_urllookupmetrics(t *testing.T) {
	viper.Reset()
	addMetrics("dd", nil)
}
