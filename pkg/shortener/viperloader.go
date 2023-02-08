// Package shortener package containing URL shortener functionality
package shortener

import (
	"fmt"
	"time"

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
		shorturl := pair2[0]
		longurl := pair2[1]
		group := pair2[2]
		timestamp := pair2[3]
		clarkezoneLog.Debugf("%s: %s %s %s \n", shorturl, longurl, group, timestamp)
		if ls != nil {
			t, err := time.Parse("2006-01-02T15:04:05-0700", timestamp.(string))
			if err != nil {
				return err
			}
			st := URLEntry{shorturl.(string), longurl.(string), group.(string), t}
			err = ls.Store(st.ShortLink, &st)
			if err != nil {
				clarkezoneLog.Debugf("ViperLoader init: Error %v", err)
				return err
			}
		} else {
			clarkezoneLog.Debugf("Lookup service is nil skipping %v", shorturl.(string))
		}
	}
	return nil
}
