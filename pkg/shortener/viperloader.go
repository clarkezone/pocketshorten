// Package shortener package containing URL shortener functionality
package shortener

import (
	"fmt"

	"github.com/spf13/viper"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

type viperLoader struct {
}

func (vl *viperLoader) Init(ls urlLookupService) error {
	values := viper.Get("values")
	if values == nil {
		clarkezoneLog.Debugf("viperLoad: no urls found in config", values)
		return fmt.Errorf("values collection not found in config json")
	}
	values2 := values.([]interface{})

	if values2 == nil {
		clarkezoneLog.Debugf("Shortenstatestore Valus is nil: %v", values)
	} else {

		clarkezoneLog.Debugf("Shortenstatestore Valus is not nil: number in collection %v", len(values2))
	}

	//Iterate over the string pairs in the array and add to lookup service
	for _, pair := range values2 {
		pair2 := pair.([]interface{})
		key := pair2[0]
		value := pair2[1]
		clarkezoneLog.Debugf("%s: %s\n", key, value)
		if ls != nil {
			err := ls.Store(key.(string), value.(string))
			if err != nil {
				clarkezoneLog.Debugf("ViperLoader init: Error %v", err)
				return err
			}
		} else {
			clarkezoneLog.Debugf("Lookup service is nil skipping %v", key.(string))
		}
	}
	return nil
}
