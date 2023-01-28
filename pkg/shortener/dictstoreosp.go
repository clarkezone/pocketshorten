// Package shortener package containing URL shortener functionality
package shortener

import (
	"errors"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

//
//lint:ignore U1000 reason backend not selected
type dictStore struct {
	m     map[string]string
	ready bool
}

// newDictStore initializes a new store backed by a dict
func newDictStore(ls storeLoader) *dictStore {
	clarkezoneLog.Debugf("NewDictStore called with loader %v", ls)
	ds := &dictStore{}
	ds.m = make(map[string]string)
	ds.ready = true
	if ls != nil {
		err := ls.Init(ds)
		if err != nil {
			ds.ready = false
		}
	}
	return ds
}

func (store *dictStore) Store(short string, long string) error {
	//TODO telemetry
	clarkezoneLog.Debugf("dictStore store short %v long %v", short, long)
	store.m[short] = long
	return nil
}

func (store *dictStore) Lookup(short string) (string, error) {
	//TODO telemetry
	val, pr := store.m[short]
	if pr {
		clarkezoneLog.Debugf("dictStore lookup short %v found %v", short, pr)
		return val, nil
	}
	clarkezoneLog.Debugf("dictstore keynotfound for %v", short)
	return "", errors.New("key not found")
}

func (store *dictStore) Count() int {
	return len(store.m)
}

func (store *dictStore) Ready() bool {
	return store.ready
}

// end dictstore
