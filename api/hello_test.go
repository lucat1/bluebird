package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()
	HelloWorld(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	msg := Message{}
	if err := json.Unmarshal(data, &msg); err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if msg.Message != "Hello World!" {
		t.Errorf("expected Hello World! got %v", msg)
	}
}
