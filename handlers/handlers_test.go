package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVersion(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	w := httptest.NewRecorder()

	GetVersion(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
	if string(data) != "v0.0.1" {
		t.Errorf("Expected v0.0.1 got %v", string(data))
	}
}
