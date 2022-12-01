package cmd

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_testserver(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	gh := getHelloHandler()
	gh(w, req)
	res := w.Result()
	//nolint
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "Hello World<BR>\n" {
		t.Errorf("expected Hello World<BR> got %v", string(data))
	}
}
