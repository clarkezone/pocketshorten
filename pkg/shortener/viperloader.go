// Package shortener package containing URL shortener functionality
package shortener

import (
	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
	"github.com/spf13/viper"
)

type viperLoader struct {
}

func (vl *viperLoader) Init(ls urlLookupService) {

	values := viper.Get("values")
	if values == nil {
		panic("no data")
	}
	values2 := values.([]interface{})

	if values2 == nil {
		clarkezoneLog.Debugf("Shortenstatestore Valus is nil: %v", values)
	} else {

		clarkezoneLog.Debugf("Shortenstatestore Valus is not nil: number in collection %v", len(values2))
	}

	// Iterate over the string pairs in the array
	//	for _, pair := range values {
	//		key := pair[0]
	//		value := pair[1]
	//		clarkezoneLog.Debugf("%s: %s\n", key, value)
	//	}
}
