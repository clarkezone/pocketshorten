package shortener

import (
	"errors"

	clarkezoneLog "github.com/clarkezone/pocketshorten/pkg/log"
)

// dictStore
//
//lint:ignore U1000 reason backend not selected
type dictStore struct {
	m map[string]string
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

// end dictstore
